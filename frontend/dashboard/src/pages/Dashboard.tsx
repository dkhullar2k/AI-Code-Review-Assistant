import { useEffect, useState } from "react";
import { api } from "../services/api";

export default function Dashboard() {
  const [reviews, setReviews] = useState([]);

  useEffect(() => {
    api.get("/reviews").then((res) => {
      setReviews(res.data);
    });
  }, []);

  const getScoreColor = (score: number) => {
    if (score >= 8) return "bg-green-100 text-green-600";
    if (score >= 5) return "bg-yellow-100 text-yellow-600";
    return "bg-red-100 text-red-600";
  };

  return (
    <div className="min-h-screen bg-gray-100 p-6">

      {/* Header */}
      <h1 className="text-3xl font-bold mb-6">
        🤖 AI Code Review Dashboard
      </h1>

      {/* Summary Cards */}
      <div className="grid grid-cols-3 gap-4 mb-6">

        <div className="bg-white p-4 rounded-xl shadow">
          <p className="text-gray-500">Total Reviews</p>
          <h2 className="text-2xl font-bold">{reviews.length}</h2>
        </div>

        <div className="bg-white p-4 rounded-xl shadow">
          <p className="text-gray-500">High Risk PRs</p>
          <h2 className="text-2xl font-bold">
            {reviews.filter((r:any) => r.score < 5).length}
          </h2>
        </div>

        <div className="bg-white p-4 rounded-xl shadow">
          <p className="text-gray-500">Avg Score</p>
          <h2 className="text-2xl font-bold">
            {(
              reviews.reduce((acc:any, r:any) => acc + (r.score || 0), 0) /
              (reviews.length || 1)
            ).toFixed(1)}
          </h2>
        </div>

      </div>

      {/* Reviews */}
      <div className="grid gap-6">

        {reviews.length === 0 && (
          <div className="text-center text-gray-500 mt-10">
            No reviews yet. Trigger a PR to see results.
          </div>
        )}

        {reviews.map((review: any, index) => (
          <div
            key={index}
            className="bg-white shadow-md rounded-xl p-5 border"
          >
            <div className="flex justify-between items-center mb-3">
              <h2 className="text-xl font-semibold">
                PR #{review.pr_id}
              </h2>

              <span className={`px-3 py-1 text-sm rounded-full ${getScoreColor(review.score)}`}>
                Score: {review.score || "N/A"}
              </span>
            </div>

            <div className="mt-3 text-sm bg-gray-50 p-3 rounded-lg">
              <pre className="whitespace-pre-wrap">
                {review.review_text}
              </pre>
            </div>
          </div>
        ))}

      </div>
    </div>
  );
}