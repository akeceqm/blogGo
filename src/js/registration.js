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

      fetch("/registration", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      })
        .then((response) => {
          if (!response.ok) {
            throw new Error("Неудачная регистрация");
          }
          return response.json();
        })
        .then((data) => {
          let login = data.login;
          let password = data.password;

          Swal.fire({
            title: "Успех!",
            html: `<div style="white-space: pre-line;">Логин: ${login}\nПароль: ${password}</div>`,
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
            text: error.message || "Неудачная регистрация",
            background: "#1e1e1e",
            color: "#fff",
          });
        });
    });
  }
});
