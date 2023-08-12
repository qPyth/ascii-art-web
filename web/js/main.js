document.getElementById("converter-form").addEventListener("submit", function(event) {
    event.preventDefault();

    var text = document.getElementById("input-text").value;
    var font = document.getElementById("font-select").value;

    var url = "ascii-art"; // Замените на фактический URL-адрес

    var data = {
        text: text,
        font: font
    };

    fetch(url, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(data)
    })
    .then(response => response.json())
    .then(responseData => {
        var outputElement = document.getElementById("output-text");
        outputElement.textContent = responseData.asciiText;
        console.log(responseData.asciiText);
    })
    .catch(error => {
        console.error('Ошибка:', error);
    });
});
