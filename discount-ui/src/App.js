import React, { useState } from "react";
import './App.css'; // Importing the CSS file for styling

const App = () => {
  const [pincode, setPincode] = useState("");
  const [result, setResult] = useState(null);

  const handleSubmit = async (e) => {
    e.preventDefault();
    const payload = {
      pincode: pincode,
    };

    try {
      const response = await fetch("http://localhost:8080/evaluate", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(payload),
      });
      const data = await response.json();
      setResult(data);
    } catch (error) {
      console.error("Error:", error);
      setResult({
        discountPercent: 0,
        message: "An error occurred while fetching the data.",
      });
    }
  };

  return (
    <div className="app-container">
      <h1 className="title">Discount Rule Engine (Pincode-Based)</h1>
      <form className="form-container" onSubmit={handleSubmit}>
        <div className="input-group">
          <label className="input-label">Pincode: </label>
          <input
            className="input-field"
            type="text"
            value={pincode}
            onChange={(e) => setPincode(e.target.value)}
          />
        </div>
        <button type="submit" className="submit-button">Check Discount</button>
      </form>

      {result && (
        <div className="result-container">
          <h2>Result:</h2>
          <p><strong>Discount:</strong> {result.discountPercent}%</p>
          <p><strong>Message:</strong> {result.message}</p>
        </div>
      )}
    </div>
  );
};

export default App;
