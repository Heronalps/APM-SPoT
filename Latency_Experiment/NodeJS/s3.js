var AWSXRay = require('aws-xray-sdk-core');
var AWS = AWSXRay.captureAWS(require('aws-sdk'));

s3 = new AWS.S3({signatureVersion: 'v4'});


exports.handler = (event, context, callback) => {


    var params = {Bucket: "latency-experiment", Key: "Michael", Body: "Hello, Michael"};

    s3.putObject(params, function(err, data) {
        if (err)
          {  console.log(err) }
        else {
          console.log('success!')
        }
    });
};
