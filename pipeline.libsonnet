{
  test:: {
    kind: 'pipeline',
    name: 'testing',
    platform: {
      os: 'linux',
      arch: 'amd64',
    },
    steps: [
      // {
      //   name: 'lint',
      //   image: 'golangci/golangci-lint:v1.46.2',
      //   pull: 'always',
      //   commands: [
      //     'golangci-lint run -v',
      //   ],
      //   volumes: [
      //     {
      //       name: 'gopath',
      //       path: '/go',
      //     },
      //   ],
      // },
      {
        name: 'embedmd',
        image: 'golang:1.18',
        pull: 'always',
        commands: [
          'make embedmd',
        ],
        volumes: [
          {
            name: 'gopath',
            path: '/go',
          },
        ],
      },
      {
        name: 'hadolint',
        image: 'hadolint/hadolint:latest-debian',
        pull: 'always',
        commands: [
          'hadolint --version',
          'hadolint docker/Dockerfile.linux.amd64',
          'hadolint docker/Dockerfile.linux.arm64',
          'hadolint docker/Dockerfile.linux.arm',
        ],
        volumes: [
          {
            name: 'gopath',
            path: '/go',
          },
        ],
      },
      // {
      //   name: 'test',
      //   image: 'golang:1.18',
      //   pull: 'always',
      //   environment: {
      //     ANDROID_API_KEY: { 'from_secret': 'android_api_key' },
      //     ANDROID_TEST_TOKEN: { 'from_secret': 'android_test_token' },
      //   },
      //   commands: [
      //     'make test',
      //   ],
      //   volumes: [
      //     {
      //       name: 'gopath',
      //       path: '/go',
      //     },
      //   ],
      // },
      // {
      //   name: 'codecov',
      //   image: 'robertstettner/drone-codecov',
      //   pull: 'always',
      //   settings: {
      //     token: { 'from_secret': 'codecov_token' },
      //   },
      // },
    ],
    volumes: [
      {
        name: 'gopath',
        temp: {},
      },
    ],
    services: [
      {
        name: 'redis',
        image: 'redis',
      },
      {
        name: 'nsq',
        image: 'nsqio/nsq',
        commands: [
          "/nsqd",
        ],
      },
    ],
  },

  build(name, os='linux', arch='amd64'):: {
    kind: 'pipeline',
    name: os + '-' + arch,
    platform: {
      os: os,
      arch: arch,
    },
    steps: [
      {
        name: 'build-push',
        image: 'golang:1.18',
        pull: 'always',
        environment: {
          CGO_ENABLED: '0',
        },
        commands: [
          'go build -v -ldflags \'-X main.build=${DRONE_BUILD_NUMBER}\' -a -o release/' + os + '/' + arch + '/' + name,
        ],
        when: {
          event: {
            exclude: [ 'tag' ],
          },
        },
      },
      // {
      //   name: 'build-push-lambda',
      //   image: 'golang:1.18',
      //   pull: 'always',
      //   environment: {
      //     CGO_ENABLED: '0',
      //   },
      //   commands: [
      //     'go build -v -tags \'lambda\' -ldflags \'-X main.build=${DRONE_BUILD_NUMBER}\' -a -o release/' + os + '/' + arch + '/lambda/' + name,
      //   ],
      //   when: {
      //     event: {
      //       exclude: [ 'tag' ],
      //     },
      //   },
      // },
      {
        name: 'build-tag',
        image: 'golang:1.18',
        pull: 'always',
        environment: {
          CGO_ENABLED: '0',
        },
        commands: [
          'go build -v -ldflags \'-X main.version=${DRONE_TAG##v} -X main.build=${DRONE_BUILD_NUMBER}\' -a -o release/' + os + '/' + arch + '/' + name,
        ],
        when: {
          event: [ 'tag' ],
        },
      },
      {
        name: 'executable',
        image: 'golang:1.18',
        pull: 'always',
        commands: [
          './release/' + os + '/' + arch + '/' + name + ' --help',
        ],
      },
      {
        name: 'publish',
        image: 'plugins/docker:' + os + '-' + arch,
        pull: 'always',
        settings: {
          daemon_off: 'false',
          auto_tag: true,
          auto_tag_suffix: os + '-' + arch,
          dockerfile: 'docker/Dockerfile.' + os + '.' + arch,
          repo: 'appleboy/' + name,
          cache_from: 'appleboy/' + name,
          username: { 'from_secret': 'docker_username' },
          password: { 'from_secret': 'docker_password' },
        },
        when: {
          event: {
            exclude: [ 'pull_request' ],
          },
        },
      },
    ],
    depends_on: [
      'testing',
    ],
    trigger: {
      ref: [
        'refs/heads/master',
        'refs/pull/**',
        'refs/tags/**',
      ],
    },
  },

  release:: {
    kind: 'pipeline',
    name: 'release-binary',
    platform: {
      os: 'linux',
      arch: 'amd64',
    },
    steps: [
      {
        name: 'build-all-binary',
        image: 'golang:1.18',
        pull: 'always',
        commands: [
          'make release'
        ],
        when: {
          event: [ 'tag' ],
        },
      },
      {
        name: 'deploy-all-binary',
        image: 'plugins/github-release',
        pull: 'always',
        settings: {
          files: [ 'dist/release/*' ],
          api_key: { 'from_secret': 'github_release_api_key' },
        },
        when: {
          event: [ 'tag' ],
        },
      },
    ],
    depends_on: [
      'testing',
    ],
    trigger: {
      ref: [
        'refs/tags/**',
      ],
    },
  },

  notifications(os='linux', arch='amd64', depends_on=[]):: {
    kind: 'pipeline',
    name: 'notifications',
    platform: {
      os: os,
      arch: arch,
    },
    steps: [
      {
        name: 'manifest',
        image: 'plugins/manifest',
        pull: 'always',
        settings: {
          username: { from_secret: 'docker_username' },
          password: { from_secret: 'docker_password' },
          spec: 'docker/manifest.tmpl',
          ignore_missing: true,
        },
      },
    ],
    depends_on: depends_on,
    trigger: {
      ref: [
        'refs/heads/master',
        'refs/tags/**',
      ],
    },
  },

  signature(key):: {
    kind: 'signature',
    hmac: key,
  },
}
