var messages = require('./gorush_pb');
var services = require('./gorush_grpc_pb');

var grpc = require('grpc');

function main() {
  var client = new services.GorushClient('localhost:9000',
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
