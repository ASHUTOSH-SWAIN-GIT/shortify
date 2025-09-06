// URL Shortener JavaScript Logic
document.addEventListener('DOMContentLoaded', function() {
  const form = document.getElementById("url-form");
  const resultDiv = document.getElementById("result");
  const shortUrlLink = document.getElementById("short-url");

  form.addEventListener("submit", async (event) => {
    event.preventDefault(); // Prevent the default form submission

    const formData = new FormData(form);

    try {
      // Send the form data to our Go backend endpoint
      const response = await fetch("/api/shorten", {
        method: "POST",
        body: new URLSearchParams(formData), // Correctly format for Go's r.FormValue
      });

      if (response.ok) {
        const data = await response.json(); // Parse JSON response
        console.log("Response from server:", data);

        // Display the shortened URL
        resultDiv.classList.remove("hidden");
        shortUrlLink.textContent = data.short_url;
        shortUrlLink.href = data.short_url;
      } else {
        alert("Failed to shorten URL. Server returned an error.");
      }
    } catch (error) {
      console.error("Error:", error);
      alert("An error occurred while contacting the server.");
    }
  });
});
    