document.addEventListener("DOMContentLoaded", function () {
    const logOutProfile = document.getElementById("logOut");
    const urlParams = new URLSearchParams(window.location.search);
    const userId = urlParams.get("userId");

    if (userId && logOutProfile) {
        logOutProfile.addEventListener("click", function (event) {
            event.preventDefault();

            fetch(`/session/delete/${userId}`, {
                method: 'DELETE',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({})
            })
                .then(response => {
                    if (!response.ok) {
                        throw new Error(`Failed to delete session: ${response.status} ${response.statusText}`);
                    }
                    return response.json();
                })
                .then(data => {
                    console.log("Session deleted successfully:", data);
                    window.location.href = "/";
                })
                .catch(error => {
                    console.error("Error:", error);
                });
        });
    }
});
