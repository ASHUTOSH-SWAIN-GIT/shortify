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
        const shortUrl = await response.text(); // In future, this will be the short URL
        console.log("Response from server:", shortUrl);

        // For now, we just show a message. In Part 2, we will display the link.
        resultDiv.classList.remove("hidden");
        shortUrlLink.textContent = "Successfully sent to server!";
        shortUrlLink.href = "#";
      } else {
        alert("Failed to shorten URL. Server returned an error.");
      }
    } catch (error) {
      console.error("Error:", error);
      alert("An error occurred while contacting the server.");
    }
  });
});
    