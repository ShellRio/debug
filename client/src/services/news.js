import axios from "axios";
import { NEWSAPI_KEY } from "../Settings";

const newsApiKey = NEWSAPI_KEY;

export function getAllNews(selectedCategory, selectedLanguage) {
  return new Promise((resolve, reject) => {
    axios
      .get("https://newsapi.org/v2/everything", {
      params:{
            q: selectedCategory,
            sortBy:"publishedAt",
            language: selectedLanguage,
            apiKey: newsApiKey
          }
        })
      .then((res) => {
        resolve(res);
      })
      .catch((err) => {
        reject(err);
      });
  })
}