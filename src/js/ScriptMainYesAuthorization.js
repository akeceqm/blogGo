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

    xhr.open('GET', `/posts/order/${UpdatePosts}`, true);

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

        clonedContainer.querySelector('.username a').textContent = item.nick_name;
        clonedContainer.querySelector('.username a').href = `/profileUser?userId=${item.author_id}`;
        if (item.Avatar.length < 6) {
            clonedContainer.querySelector('.post__avatar').src = "/assets/img/avatar.svg";
        } else {
            clonedContainer.querySelector('.post__avatar').src = item.Avatar;
        }
        clonedContainer.querySelector('.timestamp').textContent = item.DateCreatedFormat;
        clonedContainer.querySelector('.title').textContent = item.Title;
        clonedContainer.querySelector('.content').textContent = item.Text;
        clonedContainer.querySelector('.commentCount').textContent = "Комментарии: " + item.comment_count;
        clonedContainer.querySelector('.commentCount').href = `/h/post/${item.Id}/comments`;

        if (clonedContainer.querySelector('.BtnSharePost') !== null) {
            clonedContainer.querySelector('.BtnSharePost').title = item.Id;
            clonedContainer.querySelector('.BtnSharePost').id = 'BtnShare_' + item.Id;
        }
        SharePostInicialization()

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

    xhr.open('GET', `/posts/order/${UpdatePosts+1}`, true);

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
    SharePostInicialization();
};

function SharePostInicialization() {
    const btnsSharePost = document.querySelectorAll('.BtnSharePost');
    btnsSharePost.forEach((btn) => {
        btn.addEventListener('click', () => {
            let PostId = btn.title
            const currentUrl = window.location.href;
            const url = new URL(currentUrl);
            const baseUrl = `${url.protocol}//${url.hostname}${url.port ? `:${url.port}` : ''}`;

            navigator.clipboard.writeText(`${baseUrl}/h/post/${PostId}/comments`);
            alert('Ссылка скопирована в буфер обмена');
        });
    });
}

