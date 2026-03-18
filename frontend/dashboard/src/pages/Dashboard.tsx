import { useEffect, useState } from "react";
import { api } from "../services/api";

export default function Dashboard() {

  const [reviews, setReviews] = useState([]);

  useEffect(() => {
    api.get("/reviews").then((res) => {
      setReviews(res.data);
    });
  }, []);

  return (
    <div style={{ padding: 40 }}>
      <h1>AI Code Review Dashboard</h1>

      {reviews.map((review: any, index) => (
        <div key={index} style={{ marginTop: 20 }}>
          <h3>PR #{review.pr_id}</h3>
          <p>{review.review_text}</p>
        </div>
      ))}
    </div>
  );
}