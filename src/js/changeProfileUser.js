document.addEventListener('DOMContentLoaded', function() {
    const nameInput = document.getElementById('name-user');
    const avatarInput = document.getElementById('avatar');
    const descriptionInput = document.getElementById('description');
    const avatarImg = document.querySelector('.avatar');
    const updateButton = document.getElementById('update-button');

    const path = window.location.pathname;
    const userId = path.split('/').pop();
    console.log('User ID:', userId);

    let currentAvatar = '';

    // Функция для обновления профиля пользователя
    function updateProfile(data) {
        fetch(`/changeProfile/${userId}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error(`Failed to update profile: ${response.status} ${response.statusText}`);
                }
                return response.json();
            })
            .then(data => {
                console.log('Server response:', data);
                if (!data.avatar || data.avatar === '') {
                    avatarImg.src = '/assets/img/avatar.svg';
                    avatarImg.alt = 'Default Avatar';
                } else {
                    avatarImg.src = data.avatar;
                    avatarImg.alt = 'Avatar';
                }
                console.log('Redirecting to profile page');
                window.location.href = `/profileUser?userId=${userId}`;
            })
            .catch(error => {
                console.error('Failed to send data:', error.message);
                // Обработка ошибок отправки данных, например, показ сообщения пользователю
            });
    }

    // Загрузка данных пользователя при загрузке страницы
    fetch(`/api/user/${userId}`)
        .then(response => {
            if (!response.ok) {
                throw new Error('Failed to fetch user data');
            }
            return response.json();
        })
        .then(data => {
            if (!data.avatar || data.avatar === '') {
                avatarImg.src = '/assets/img/avatar.svg';
                avatarImg.alt = 'Default Avatar';
            } else {
                avatarImg.src = data.avatar;
                avatarImg.alt = 'Avatar';
                currentAvatar = data.avatar; // Сохраняем текущий URL аватара
            }

            avatarImg.onerror = function() {
                avatarImg.src = '/assets/img/avatar.svg';
                avatarImg.alt = 'Default Avatar';
            };

            nameInput.value = data.nick_name || '';
            descriptionInput.value = data.description || '';
            console.log(data);
        })
        .catch(error => {
            console.error('Failed to load user data', error);
        });

    // Обработка изменений в поле выбора аватара
    avatarInput.addEventListener('change', function() {
        const avatarFile = avatarInput.files[0];
        if (!avatarFile) {
            return;
        }

        const reader = new FileReader();
        reader.onload = function(e) {
            avatarImg.src = e.target.result;
            avatarImg.alt = 'New Avatar Preview';
        };
        reader.readAsDataURL(avatarFile);
    });

    // Обработка клика по кнопке обновления профиля
    updateButton.addEventListener('click', function() {
        console.log('Update button clicked');
        const newName = nameInput.value.trim();
        const newDescription = descriptionInput.value.trim();
        const avatarFile = avatarInput.files[0];

        let data = {
            nick_name: newName,
            description: newDescription
        };

        if (avatarFile) {
            const reader = new FileReader();
            reader.onload = function(e) {
                const avatarBase64 = e.target.result.split(',')[1]; // Получаем только base64 часть
                data.avatar = avatarBase64;

                updateProfile(data);
            };
            reader.readAsDataURL(avatarFile);
        } else {
            // Если аватар не выбран заново и текущий URL аватара есть, отправляем его
            if (currentAvatar) {
                data.avatar = currentAvatar;
                updateProfile(data); // Отправляем данные на обновление профиля
            } else {
                // Если аватар не выбран и текущего URL аватара нет, обновляем без аватара
                updateProfile(data); // Отправляем данные на обновление профиля
            }
        }

        console.log('New Name:', newName);
        console.log('New Description:', newDescription);
        console.log('Data to send:', data);
    });
});
