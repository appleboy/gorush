/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

var messages = require('./gorush_pb');
var services = require('./gorush_grpc_pb');

var grpc = require('grpc');

function main() {
  var client = new services.GorushClient('localhost:50051',
    grpc.credentials.createInsecure());
  var request = new messages.NotificationRequest();
  request.setPlatform(2);
  request.setTokensList(["1234567890"]);
  request.setMessage("Hello!!");
  client.send(request, function (err, response) {
    console.log('Success:', response.getSuccess());
    console.log('Counts:', response.getCounts());
  });
}

main();
