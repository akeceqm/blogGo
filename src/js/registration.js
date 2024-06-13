document.addEventListener("DOMContentLoaded", function () {
  let form = document.getElementById("registrationForm");

  if (form) {
    form.addEventListener("submit", function (e) {
      e.preventDefault();

      let nameInput = document.querySelector("input[name='nick_name']");
      let emailInput = document.querySelector("input[name='email']");

      let data = {
        nick_name: nameInput.value,
        email: emailInput.value,
      };

      let xhr = new XMLHttpRequest();

      xhr.open("POST", "/registration", true);
      xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");

      xhr.onload = function () {
        if (xhr.status >= 200 && xhr.status < 300) {
          let response = JSON.parse(xhr.responseText);
          let login = response.login;
          let password = response.password;

          Swal.fire({
            title: "Успех!",
            html: `<div style="white-space: pre-line;">Логин: ${login}\nПароль: ${password}</div>`,
            icon: "success",
            background: "#1e1e1e",
            color: "#fff",
          });
        } else {
          console.error("Error:", xhr.statusText);
          Swal.fire({
            icon: "error",
            title: "Ошибка!",
            text: "Неудачная регистрация",
            background: "#1e1e1e",
            color: "#fff",
          });
        }
      };

      xhr.onerror = function () {
        console.error("Request failed");
        Swal.fire({
          title: "Интернет?",
          text: "Проблемы с сервером либо интернетом",
          icon: "question",
          background: "#1e1e1e",
          color: "#fff",
        });
      };

      xhr.send(JSON.stringify(data));
    });
  }
});
