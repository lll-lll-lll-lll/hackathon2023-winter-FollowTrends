import * as functions from "firebase-functions";

// Start writing Firebase Functions
// https://firebase.google.com/docs/functions/typescript

// export const helloWorld = functions.https.onRequest((request, response) => {
//   functions.logger.info("Hello logs!", {structuredData: true});
//   response.send("Hello from Firebase!");
// });

export const notify = functions.region("asia-northeast1")
    .pubsub.schedule("every 10 minutes")
    .timeZone("Asia/Tokyo")
    .onRun((context) => {
      console.log("Hello World!");
    });
