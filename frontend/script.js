// Fetch data based on year selection
document.getElementById("year-select").addEventListener("change", fetchData);

function fetchData() {
    const selectedYear = document.getElementById("year-select").value;
    const apiUrl = `http://localhost:8080/api/year/${selectedYear}`;

    fetch(apiUrl)
        .then(response => {
            if (!response.ok) {
                throw new Error("Network response was not ok");
            }
            return response.json();
        })
        .then(data => {
            // Update welcome message
            document.getElementById("welcome-message").innerText = data.welcomeMessage;

            // Update Secretary
            updatePerson(selectedYear, "secretary", data.secretary);

            // Update Joint Secretaries
            updatePerson(selectedYear, "joint-secy1", data.jointSecretaries[0]);
            updatePerson(selectedYear, "joint-secy2", data.jointSecretaries[1]);
        })
        .catch(error => console.error("Error fetching data:", error));
}

function updatePerson(selectedYear, prefix, person) {
    // Update the image by sending a request for the image
    const apiUrl = `http://localhost:8080/${person.photo}`
    fetch(apiUrl)
        .then(response => {
            if (!response.ok) {
                throw new Error(`Error fetching image: ${person.photo}`);
            }
            return response.blob();  // Convert the image response to blob
        })
        .then(imageBlob => {
            const imageUrl = URL.createObjectURL(imageBlob);
            document.getElementById(`${prefix}-photo`).src = imageUrl;
        })
        .catch(error => console.error("Error fetching image:", error));

    // Update the name and one-liner
    document.getElementById(`${prefix}-name`).innerText = person.name;
    document.getElementById(`${prefix}-one-liner`).innerText = person.oneLiner;
}

fetchData()