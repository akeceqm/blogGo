const openPopUp = document.getElementById('BtnCreateNewPost')
const closePopUp = document.getElementById('BtnClose')
const popUp = document.getElementById('Pop_up')
const BtnCreatePost = document.getElementById('BtnSucces')
const BtnEditPost = document.getElementById('BtnSuccesEdit')
const BtnDelPost = document.getElementById('BtnSuccesDel')
AddEditAndDelete()

openPopUp.onclick = () => {
    popUp.classList.add('active')
    document.addEventListener("keydown", handleEnterPress)
}
closePopUp.addEventListener('click', () => {
    popUp.classList.remove('active')
    document.removeEventListener("keydown", handleEnterPress)
})

function AddEditAndDelete() {
    const btnsEditPost = document.querySelectorAll('.BtnEditPost');
    const btnCloseEdit = document.getElementById('BtnCloseEdit')
    const popUpEdit = document.getElementById('Pop_upEdit')

    btnCloseEdit.onclick = () => {
        popUpEdit.classList.remove('active')
        document.removeEventListener("keydown", handleEnterPressEdit)
    }
    btnsEditPost.forEach((btn) => {
        btn.addEventListener('click', () => {
            popUpEdit.classList.add('active')
            document.addEventListener("keydown", handleEnterPressEdit)

            const TitleInput = document.getElementById('TitleInputEdit')
            const TextInput = document.getElementById('TextInputEdit')
            let TitleAndText = btn.title.split('|')
            TitleInput.value = TitleAndText[0]
            TextInput.value = TitleAndText[1]
            TitleInput.title = btn.id.replace('BtnEdit_', '')
        });
    });

    const btnsDelPost = document.querySelectorAll('.BtnDelPost');
    const btnCloseDel = document.getElementById('BtnCloseDel')
    const popUpDel = document.getElementById('Pop_upDel')

    btnCloseDel.onclick = () => {
        popUpDel.classList.remove('active')
        document.removeEventListener("keydown", handleEnterPressDel)
    }
    btnsDelPost.forEach((btn) => {
        btn.addEventListener('click', () => {
            popUpDel.classList.add('active')
            document.addEventListener("keydown", handleEnterPressDel)

            const TitleInput = document.getElementById('TitleInputDel')
            const TextInput = document.getElementById('TextInputDel')
            let TitleAndText = btn.title.split('|')
            TitleInput.value = TitleAndText[0]
            TextInput.value = TitleAndText[1]
            TitleInput.title = btn.id.replace('BtnDel_', '')
        });
    });

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

function handleEnterPress(event) {
    if (event.key === 'Enter') {
        createNewPost();
    }
}
function handleEnterPressEdit(event) {
    if (event.key === 'Enter') {
        editPost();
    }
}
function handleEnterPressDel(event) {
    if (event.key === 'Enter') {
        deletePost();
    }
}

BtnCreatePost.onclick = createNewPost;
BtnEditPost.onclick = editPost;
BtnDelPost.onclick = deletePost;

function createNewPost() {
    const TitleInput = document.getElementById('TitleInput')
    const TextInput = document.getElementById('TextInput')
    const currentUrl = window.location.href;
    const url = new URL(currentUrl);
    const userId = url.searchParams.get('userId');

    if (TitleInput.value == "") {
        TitleInput.style.borderColor = '#ff0000'
        return
    }else {
        TitleInput.style.borderColor = 'transparent'
    }
    if (TextInput.value == "") {
        TextInput.style.borderColor = '#ff0000'
        return;
    }else {
        TextInput.style.borderColor = 'transparent'
    }

    let xhr = new XMLHttpRequest();
    xhr.open("POST", `/posts`);
    xhr.onload = function () {
        if (xhr.status >= 200 && xhr.status < 300) {
            let response = JSON.parse(xhr.responseText);
            console.log(response);
            location.reload();
        } else {
            console.error("Server error:", xhr.status, xhr.statusText);
            location.reload();
        }
    };


    let newPostData = JSON.stringify({
        title: TitleInput.value,
        text: TextInput.value,
        author_id: userId,
    });
    console.log(newPostData);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.send(newPostData);
    TitleInput.value = ""
    TextInput.value = ""
}

function editPost() {
    const TitleInputEdit = document.getElementById('TitleInputEdit')
    const TextInputEdit = document.getElementById('TextInputEdit')

    if (TitleInputEdit.value == "") {
        TitleInputEdit.style.borderColor = '#ff0000'
        return
    }else {
        TitleInputEdit.style.borderColor = 'transparent'
    }
    if (TextInputEdit.value == "") {
        TextInputEdit.style.borderColor = '#ff0000'
        return;
    }else {
        TextInputEdit.style.borderColor = 'transparent'
    }

    let xhr = new XMLHttpRequest();
    xhr.open("PUT", `/posts`);
    xhr.onload = function () {
        if (xhr.status >= 200 && xhr.status < 300) {
            let response = JSON.parse(xhr.responseText);
            console.log(response);
            location.reload();
        } else {
            console.error("Server error:", xhr.status, xhr.statusText);
            location.reload();
        }
    };

    let PostData = JSON.stringify({
        id: TitleInputEdit.title,
        title: TitleInputEdit.value,
        text: TextInputEdit.value,
    });
    console.log(PostData);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.send(PostData);
    TitleInputEdit.value = ""
    TextInputEdit.value = ""
}

function deletePost() {
    const TitleInputDel = document.getElementById('TitleInputDel')
    const TextInputDel = document.getElementById('TextInputDel')

    let xhr = new XMLHttpRequest();
    xhr.open("DELETE", `/posts/${TitleInputDel.title}`);
    xhr.onload = function () {
        if (xhr.status >= 200 && xhr.status < 300) {
            let response = JSON.parse(xhr.responseText);
            console.log(response);
            location.reload();
        } else {
            console.error("Server error:", xhr.status, xhr.statusText);
            location.reload();
        }
    };
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.send();
    TitleInputDel.value = ""
    TextInputDel.value = ""
}


const btnLoadPosts = document.getElementById("BtnLoadPosts");
if (btnLoadPosts) {
    btnLoadPosts.onclick = function () {
        UpdatePosts++;
        fetchData();
    };
}

let UpdatePosts = 2;
const currentUrl = window.location.href;
const url = new URL(currentUrl);
const userId = url.searchParams.get('userId');

function fetchData() {
    var xhr = new XMLHttpRequest();

    xhr.open('GET', `/posts/order/${UpdatePosts}/${userId}`, true);

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
        if (item.Avatar.length <= 6) {
            clonedContainer.querySelector('.post__avatar').src = "/assets/img/avatar.svg";
        } else {
            clonedContainer.querySelector('.post__avatar').src = item.Avatar;
        }
        clonedContainer.querySelector('.timestamp').textContent = item.DateCreatedFormat;
        clonedContainer.querySelector('.title').textContent = item.Title;
        clonedContainer.querySelector('.content').textContent = item.Text;
        clonedContainer.querySelector('.commentCount').textContent = "Комментарии: " + item.comment_count;
        clonedContainer.querySelector('.commentCount').href = `/h/post/${item.Id}/comments`;

        if (clonedContainer.querySelector('.BtnEditPost') !== null) {
            clonedContainer.querySelector('.BtnEditPost').title = item.Title + "|" + item.Text;
            clonedContainer.querySelector('.BtnEditPost').id = 'BtnEdit_' + item.Id;
        }

        if (clonedContainer.querySelector('.BtnDelPost') !== null) {
            clonedContainer.querySelector('.BtnDelPost').title = item.Title + "|" + item.Text;
            clonedContainer.querySelector('.BtnDelPost').id = 'BtnDel_' + item.Id;
        }
        if (clonedContainer.querySelector('.BtnSharePost') !== null) {
            clonedContainer.querySelector('.BtnSharePost').title = item.Id;
            clonedContainer.querySelector('.BtnSharePost').id = 'BtnShare_' + item.Id;
        }

        dataContainer.appendChild(clonedContainer);
    });
    AddEditAndDelete()
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
    xhr.open('GET', `/posts/order/${UpdatePosts+1}/${userId}`, true);
    xhr.onreadystatechange = function () {
        if (xhr.readyState === XMLHttpRequest.DONE) {
            if (xhr.status === 200) {
                var data = JSON.parse(xhr.responseText);
                if (data === null) {
                    callback(false);
                    console.log("запрос не удался или данные равны null");
                    return;
                }
                callback(true);
                console.log("посты загружены");
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
            console.log("больше не постов");
        }
    });
};