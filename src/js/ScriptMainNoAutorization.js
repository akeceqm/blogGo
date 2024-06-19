const btnLoadPosts = document.getElementById("BtnLoadPosts");
if (btnLoadPosts) {
    btnLoadPosts.onclick = function () {
        UpdatePosts++;
        fetchData();
    };
}

let UpdatePosts = 1;

// Функция для выполнения запроса и обработки данных
function fetchData() {
    // Создаем новый объект XMLHttpRequest
    var xhr = new XMLHttpRequest();

    // Указываем метод и URL для запроса
    xhr.open('GET', `http://localhost:8080/posts/order/${UpdatePosts}`, true);

    // Определяем функцию, которая будет выполнена при изменении состояния запроса
    xhr.onreadystatechange = function () {
        // Проверяем, что запрос завершен и статус ответа 200 (OK)
        if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            // Парсим полученные данные из JSON
            var data = JSON.parse(xhr.responseText);

            // Добавляем элементы на страницу на основе полученных данных
            updatePage(data);
        }
    };

    // Отправляем запрос
    xhr.send();
}

// Функция для обновления страницы на основе данных
function updatePage(data) {
    // Находим шаблонный контейнер и основной контейнер
    var templateContainer = document.getElementById('mainPost__inner');
    var dataContainer = document.getElementById('MainPosts');
    templateContainer.removeAttribute('id');

    // Вырезаем кнопку из страницы, сохраняя ссылку на неё
    const parent = btnLoadPosts.parentNode;
    const removedButton = parent.removeChild(btnLoadPosts);

    // Проходим по массиву данных и создаем элементы
    data.forEach(function(item) {
        // Клонируем шаблонный контейнер
        var clonedContainer = templateContainer.cloneNode(true);
        //clonedContainer.style.display = 'block';

        // Заполняем клонированный контейнер данными
        clonedContainer.querySelector('.username').textContent = item.nick_name; // Используем данные из item
        clonedContainer.querySelector('.timestamp').textContent = item.DateCreatedFormat; // Используем данные из item
        clonedContainer.querySelector('.title').textContent = item.Title; // Используем данные из item
        clonedContainer.querySelector('.content').textContent = item.Text; // Используем данные из item
        clonedContainer.querySelector('.commentCount').textContent = "Комментарии: " + item.comment_count; // Используем данные из item

        // Добавляем клонированный контейнер в основной контейнер
        dataContainer.appendChild(clonedContainer);
    });
    dataContainer.appendChild(removedButton)
}

