const btnLoadPosts = document.getElementById("BtnLoadPosts");
if (btnLoadPosts) {
    btnLoadPosts.onclick = function () {
        UpdatePosts++;
        fetchData();
    };
}

let UpdatePosts = 1;

function fetchData() {
    var xhr = new XMLHttpRequest();

    const currentUrl = window.location.href;
    const url = new URL(currentUrl);
    const userId = url.searchParams.get('userId');
    console.log(userId);

    xhr.open('GET', `http://localhost:8080/posts/order/${UpdatePosts}/${userId}`, true);

    xhr.onreadystatechange = function () {
        if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            var data = JSON.parse(xhr.responseText);
            updatePage(data);
        }
    };
    xhr.send();
}

function updatePage(data) {
    var templateContainer = document.getElementById('mainPost__inner');
    var dataContainer = document.getElementById('MainPosts');
    templateContainer.removeAttribute('id');

    const MainBtnLoadPosts = document.getElementById('Main__BtnLoadPosts');
    const parent = MainBtnLoadPosts.parentNode;
    const removedButton = parent.removeChild(MainBtnLoadPosts);

    data.forEach(function(item) {
        var clonedContainer = templateContainer.cloneNode(true);

        clonedContainer.querySelector('.username').textContent = item.nick_name; // Используем данные из item
        clonedContainer.querySelector('.timestamp').textContent = item.DateCreatedFormat; // Используем данные из item
        clonedContainer.querySelector('.title').textContent = item.Title; // Используем данные из item
        clonedContainer.querySelector('.content').textContent = item.Text; // Используем данные из item
        clonedContainer.querySelector('.commentCount').textContent = "Комментарии: " + item.comment_count; // Используем данные из item

        dataContainer.appendChild(clonedContainer);
    });
    CheckLoadPosts(function(success) {
        if (!success) {
            console.log('Запрос не удался или данные равны null');
        } else {
            dataContainer.appendChild(MainBtnLoadPosts)
        }
    });
}

function CheckLoadPosts(callback) {
    var xhr = new XMLHttpRequest();

    xhr.open('GET', `http://localhost:8080/posts/order/${UpdatePosts+1}`, true);

    xhr.onreadystatechange = function () {
        if (xhr.readyState === XMLHttpRequest.DONE) {
            if (xhr.status === 200) {
                var data = JSON.parse(xhr.responseText);

                if (data === null) {
                    callback(false);
                    return;
                }
                callback(true);
            } else {
                callback(false);
            }
        }
    };
    xhr.send();
}

window.onload = function() {
    UpdatePosts = 1;
    console.log("Переменная UpdatePosts инициализирована");
    CheckLoadPosts(function(success) {
        if (!success) {
            const MainBtnLoadPosts = document.getElementById('Main__BtnLoadPosts');
            const parent = MainBtnLoadPosts.parentNode;
            const removedButton = parent.removeChild(MainBtnLoadPosts);
        }
    });
};

