const cities = [
    "Mumbai", "Delhi", "Bangalore", "Hyderabad", "Ahmedabad", "Chennai",
    "Kolkata", "Surat", "Pune", "Jaipur", "Lucknow", "Kanpur", "Nagpur",
    "Indore", "Thane", "Bhopal", "Visakhapatnam", "Patna", "Vadodara", "Ghaziabad",
    "Ludhiana", "Agra", "Nashik", "Faridabad", "Meerut", "Rajkot", "Varanasi",
    "Srinagar", "Aurangabad", "Dhanbad", "Amritsar", "Allahabad", "Ranchi",
    "Howrah", "Coimbatore", "Jabalpur", "Gwalior", "Vijayawada", "Jodhpur",
    "Madurai", "Raipur", "Kota", "Chandigarh", "Guwahati", "Solapur"
];

function populateCityDropdown(id) {
    const dropdown = document.getElementById(id);
    cities.forEach(city => {
        const option = document.createElement("option");
        option.value = city;
        option.textContent = city;
        dropdown.appendChild(option);
    });
}

document.getElementById("bookingForm").addEventListener("submit", function (event) {
    event.preventDefault();

    const firstName = document.getElementById("firstName").value;
    const lastName = document.getElementById("lastName").value;
    const email = document.getElementById("email").value;
    const source = document.getElementById("source").value;
    const destination = document.getElementById("destination").value;
    const date = document.getElementById("date").value;
    const tickets = document.getElementById("tickets").value;

    if (source === destination) {
        document.getElementById("response").textContent = "Source and destination cannot be the same.";
        return;
    }

    document.getElementById("response").textContent =
        `Booking successful! ${tickets} ticket(s) from ${source} to ${destination} on ${date} booked by ${firstName} ${lastName}.`;

        const responseElement = document.getElementById("response");
        responseElement.classList.add('show-message');

        setTimeout(() => {
            responseElement.classList.remove('show-message');
        }, 3000);

        const submitButton = document.querySelector('form button');
        submitButton.classList.add('clicked');
    
        setTimeout(() => {
          submitButton.classList.remove('clicked');
        }, 200);
});

window.onload = function () {
    const formContainer = document.querySelector('.form-container');
    formContainer.style.opacity = 0;

    setTimeout(() => {
      formContainer.style.opacity = 1;
      formContainer.style.transition = 'opacity 0.5s ease-in-out';
    }, 200);

    populateCityDropdown("source");
    populateCityDropdown("destination");
};