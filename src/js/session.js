document.getElementById('profile-link').addEventListener('click', function(event) {
    event.preventDefault();
    const sessionId = getCookie('session_id');
    if (sessionId) {
        fetch('/profile', {
            method: 'GET',
            credentials: 'include'
        }).then(response => {
            if (response.ok) {
                window.location.href = '/profile';
            } else {
                alert('Вы не авторизованы');
            }
        }).catch(error => {
            console.error('Ошибка:', error);
        });
    } else {
        alert('Кука session_id не найдена');
    }
});

function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
}