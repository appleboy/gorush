var messages = require('./gorush_pb');
var services = require('./gorush_grpc_pb');

var grpc = require('grpc');

function main() {
  var client = new services.GorushClient('localhost:9000',
    grpc.credentials.createInsecure());
  var request = new messages.NotificationRequest();
  var alert = new messages.Alert();
  request.setPlatform(2);
  request.setTokensList(["1234567890"]);
  request.setMessage("Hello!!");
  request.setTitle("hello2");
  request.setBadge(2);
  request.setCategory("mycategory");
  request.setSound("sound")
  alert.setTitle("title");
  request.setAlert(alert);
  request.setThreadid("threadID");
  request.setContentavailable(false);
  request.setMutablecontent(false);
  client.send(request, function (err, response) {
    if(err) {
      console.log(err);
    } else {
      console.log("Success:", response.getSuccess());
      console.log("Counts:", response.getCounts());
    }
  });
}

main();
