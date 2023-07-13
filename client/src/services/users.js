import axios from "axios";

export function getAll() {
  return new Promise((resolve, reject) => {
    axios
      .get("http://localhost:8080/victime/all?auth=password")
      .then((res) => {
        resolve(res.data);
      })
      .catch((err) => {
        reject(err);
      });
  });
}

export function getById(id) {
  return new Promise((resolve, reject) => {
    axios
      .get(`http://localhost:8080/victime?auth=password&id=${id}`)
      .then((res) => {
        resolve(res.data);
      })
      .catch((err) => {
        reject(err);
      });
  });
}

export function updateStatus(id, code) {
  return new Promise((resolve, reject) => {
    const url = "http://localhost:8080/status/update";
    const formData = new FormData();
    formData.append("auth", "password");
    formData.append("id", id);
    formData.append("status", code);

    axios
      .post(url, formData)
      .then((res) => {
        resolve(res.data);
      })
      .catch((err) => {
        reject(err);
      });
  });
}


