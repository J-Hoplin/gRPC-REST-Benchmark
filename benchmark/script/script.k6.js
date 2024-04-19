import { sleep } from "k6";
import http from "k6/http";

export const options = {
  vus: 100,
  duration: "10s",
};

export default function () {
  http.get(__ENV.ENDPOINT);
  sleep(1);
}
