// Minimal URL Shortener JavaScript
document.addEventListener('DOMContentLoaded', function() {
    const form = document.getElementById("url-form");
    const resultDiv = document.getElementById("result");
    const shortUrlLink = document.getElementById("short-url");
    const errorDiv = document.getElementById("error");
    const errorMessage = document.getElementById("error-message");
    const submitBtn = document.getElementById("submit-btn");
    const btnText = document.querySelector(".btn-text");
    const loading = document.getElementById("loading");
    const copyBtn = document.getElementById("copy-btn");
    const urlInput = document.getElementById("url-input");

    // Hide error and result initially
    hideError();
    hideResult();

    form.addEventListener("submit", async (event) => {
        event.preventDefault();
        
        const formData = new FormData(form);
        const url = formData.get("url");

        if (!url) {
            showError("Please enter a valid URL");
            return;
        }

        // Validate URL format
        try {
            new URL(url);
        } catch {
            showError("Please enter a valid URL format (e.g., https://example.com)");
            return;
        }

        // Show loading state
        setLoadingState(true);
        hideError();
        hideResult();

        try {
            const response = await fetch("/api/shorten", {
                method: "POST",
                body: new URLSearchParams(formData),
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                }
            });

            if (response.ok) {
                const data = await response.json();
                console.log("Response from server:", data);
                
                // Display the result
                showResult(data);
                
            } else {
                const errorText = await response.text();
                console.error("Server error:", errorText);
                showError("Failed to shorten URL. Please try again.");
            }
        } catch (error) {
            console.error("Network error:", error);
            showError("Network error. Please check your connection and try again.");
        } finally {
            setLoadingState(false);
        }
    });

    // Copy to clipboard functionality
    copyBtn.addEventListener("click", async () => {
        const shortUrl = shortUrlLink.href;
        
        try {
            await navigator.clipboard.writeText(shortUrl);
            copyBtn.textContent = "Copied!";
            copyBtn.classList.add("copied");
            
            setTimeout(() => {
                copyBtn.textContent = "Copy";
                copyBtn.classList.remove("copied");
            }, 2000);
        } catch (err) {
            console.error("Failed to copy:", err);
            // Fallback for older browsers
            const textArea = document.createElement("textarea");
            textArea.value = shortUrl;
            document.body.appendChild(textArea);
            textArea.select();
            document.execCommand("copy");
            document.body.removeChild(textArea);
            
            copyBtn.textContent = "Copied!";
            copyBtn.classList.add("copied");
            
            setTimeout(() => {
                copyBtn.textContent = "Copy";
                copyBtn.classList.remove("copied");
            }, 2000);
        }
    });

    // Auto-focus on input
    urlInput.focus();

    // Clear result when user starts typing
    urlInput.addEventListener("input", () => {
        if (resultDiv.classList.contains("show")) {
            hideResult();
        }
        if (errorDiv.classList.contains("show")) {
            hideError();
        }
    });

    function setLoadingState(isLoading) {
        if (isLoading) {
            submitBtn.disabled = true;
            btnText.style.display = "none";
            loading.style.display = "flex";
            urlInput.disabled = true;
        } else {
            submitBtn.disabled = false;
            btnText.style.display = "block";
            loading.style.display = "none";
            urlInput.disabled = false;
        }
    }

    function showResult(data) {
        shortUrlLink.textContent = data.short_url;
        shortUrlLink.href = data.short_url;
        resultDiv.classList.add("show");
    }

    function hideResult() {
        resultDiv.classList.remove("show");
    }

    function showError(message) {
        errorMessage.textContent = message;
        errorDiv.classList.add("show");
    }

    function hideError() {
        errorDiv.classList.remove("show");
    }
});