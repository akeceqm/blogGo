document.addEventListener("DOMContentLoaded", function () {
  let form = document.getElementById("authorizationForm");

  if (form) {
    form.addEventListener("submit", function (e) {
      e.preventDefault();

      let loginInput = document.querySelector("input[name='login']");
      let passwordInput = document.querySelector("input[name='password_hash']");

      let loginError = document.getElementById("loginError");
      let passwordError = document.getElementById("passwordError");

      loginError.textContent = "";
      passwordError.textContent = "";

      if (loginInput.value.trim() === "") {
        loginError.textContent = "Введите логин";
        return;
      }

      if (passwordInput.value.trim() === "") {
        passwordError.textContent = "Введите пароль";
        return;
      }

      let data = {
        login: loginInput.value,
        password_hash: passwordInput.value,
      };

      fetch("/authorization", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      })
        .then((response) => {
          if (!response.ok) {
            throw new Error("Неудачная авторизация");
          }
          return response.json();
        })
        .then((data) => {
          Swal.fire({
            title: "Успех!",
            text: "Успешная авторизация",
            icon: "success",
            background: "#1e1e1e",
            color: "#fff",
          }).then(() => {
            window.location.href = "/profileUser?userId=" + data.id;
          });
        })
        .catch((error) => {
          console.error("Ошибка:", error);
          Swal.fire({
            icon: "error",
            title: "Ошибка!",
            text: error.message || "Неудачная авторизация",
            background: "#1e1e1e",
            color: "#fff",
          });
        });
    });
  }
});
