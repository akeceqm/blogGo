const openPopUp = document.getElementById('BtnCreateNewPost')
const closePopUp = document.getElementById('BtnClose')
const popUp = document.getElementById('Pop_up')
const BtnCreatePost = document.getElementById('BtnSucces')

const BtnEditPost = document.getElementById('BtnSuccesEdit')
const btnsEditPost = document.querySelectorAll('.BtnEditPost');
const btnCloseEdit = document.getElementById('BtnCloseEdit')
const popUpEdit = document.getElementById('Pop_upEdit')

const BtnDelPost = document.getElementById('BtnSuccesDel')
const btnsDelPost = document.querySelectorAll('.BtnDelPost');
const btnCloseDel = document.getElementById('BtnCloseDel')
const popUpDel = document.getElementById('Pop_upDel')

openPopUp.onclick = () => {
    popUp.classList.add('active')
    document.addEventListener("keydown", handleEnterPress)
}
closePopUp.addEventListener('click', () => {
    popUp.classList.remove('active')
    document.removeEventListener("keydown", handleEnterPress)
})

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