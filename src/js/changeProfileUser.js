document.addEventListener('DOMContentLoaded', function() {
    const profileForm = document.getElementById('profileForm');
    const nameInput = document.getElementById('name-user');
    const avatarInput = document.getElementById('avatar');
    const descriptionInput = document.getElementById('description');
    const updateButton = document.getElementById('update-button');
    const avatarImg = document.querySelector('.avatar');

    // Получаем текущий путь URL
    const path = window.location.pathname;
    // Извлекаем userId из пути URL
    const userId = path.split('/').pop();
    console.log('User ID:', userId);

    // Загрузка данных пользователя
    fetch(`/api/user/${userId}`)
        .then(response => {
            if (!response.ok) {
                throw new Error('Failed to fetch user data');
            }
            return response.json();
        })
        .then(data => {
            // Отображение текущего аватара или аватара по умолчанию
            if (!data.avatar || data.avatar === '') {
                avatarImg.src = '/assets/img/avatar.svg';
                avatarImg.alt = 'Default Avatar';
            } else {
                avatarImg.src = data.avatar;
                avatarImg.alt = 'Avatar';
            }

            avatarImg.onerror = function() {
                avatarImg.src = '/assets/img/avatar.svg';
                avatarImg.alt = 'Default Avatar';
            };

            // Заполнение полей текущими данными пользователя
            nameInput.value = data.nick_name || '';
            descriptionInput.value = data.description || '';

            console.log(data); // Отладочная информация
        })
        .catch(error => {
            console.error('Failed to load user data', error);
        });

    // Обработчик изменения input для загрузки аватара
    avatarInput.addEventListener('change', function() {
        const avatarFile = avatarInput.files[0];

        if (!avatarFile) {
            return;
        }

        // Обновляем изображение предварительного просмотра
        const reader = new FileReader();
        reader.onload = function(e) {
            avatarImg.src = e.target.result;
            avatarImg.alt = 'New Avatar Preview';
        };
        reader.readAsDataURL(avatarFile);
    });

    // Обработчик кнопки обновления профиля
    updateButton.addEventListener('click', function() {
        const newName = nameInput.value.trim(); // Убираем лишние пробелы
        const newDescription = descriptionInput.value.trim(); // Убираем лишние пробелы
        const avatarFile = avatarInput.files[0]; // Получаем выбранный файл аватара
        const formData = new FormData();
        formData.append('nick_name', newName); // Используйте 'nick_name', а не 'name'
        formData.append('description', newDescription); // Используйте 'description', а не 'desc' в структуре

        // Если выбран файл аватара, добавьте его в formData
        if (avatarFile) {
            formData.append('avatar', avatarFile);
        }

        // Отправляем данные на сервер
        fetch(`/changeProfile/${userId}`, {
            method: 'PUT',
            body: formData
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Failed to update profile');
                }
                return response.json();
            })
            .then(data => {
                console.log('Server response:', data);
                // Обновление отображения данных на странице
                if (!data.avatar || data.avatar === '') {
                    avatarImg.src = '/assets/img/avatar.svg';
                    avatarImg.alt = 'Default Avatar';
                } else {
                    avatarImg.src = data.avatar;
                    avatarImg.alt = 'Avatar';
                }
            })
            .catch(error => {
                console.error('Failed to send data:', error);
            });
        console.log('New Name:', newName);
        console.log('New Description:', newDescription);
    });

});
