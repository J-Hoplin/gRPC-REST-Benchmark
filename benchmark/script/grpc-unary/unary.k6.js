import { sleep } from "k6";
import http from "k6/http";

export const options = {
  vus: 100,
  duration: "10s",
};

export default function () {
  http.get("http://localhost:8080/grpc/unary?from=1&to=100");
  sleep(1);
}
