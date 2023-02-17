import * as functions from "firebase-functions";
import axios from "axios";

// Start writing Firebase Functions
// https://firebase.google.com/docs/functions/typescript

export const helloWorld = functions.https.onRequest((request, response) => {
  functions.logger.info("Hello logs!", {structuredData: true});
  response.send("Hello from Firebase!");
});

export const notify = functions.region("asia-northeast1")
    .pubsub.schedule("every day 18:00")
    .timeZone("Asia/Tokyo")
    .onRun(async () => {
      console.log("notify");
      // localhost:5004/notify をここで叩く
      const res = await axios.get("https://930e-221-247-249-83.jp.ngrok.io/notify");
      console.log(res.data);
    });
