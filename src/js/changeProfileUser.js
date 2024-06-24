document.addEventListener('DOMContentLoaded', function() {
    const nameInput = document.getElementById('name-user');
    const avatarInput = document.getElementById('avatar');
    const descriptionInput = document.getElementById('description');
    const avatarImg = document.querySelector('.avatar');
    const updateButton = document.getElementById('update-button');

    const path = window.location.pathname;
    const userId = path.split('/').pop();
    console.log('User ID:', userId);

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

    updateButton.addEventListener('click', function() {
        const newName = nameInput.value.trim();
        const newDescription = descriptionInput.value.trim();
        const avatarFile = avatarInput.files[0];

        if (avatarFile) {
            const reader = new FileReader();
            reader.onload = function(e) {
                const avatarBase64 = e.target.result.split(',')[1];

                const data = {
                    nick_name: newName,
                    description: newDescription,
                    avatar: avatarBase64
                };

                fetch(`/changeProfile/${userId}`, {
                    method: 'PUT',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(data)
                })
                    .then(response => {
                        if (!response.ok) {
                            throw new Error('Failed to update profile');
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
                    })
                    .catch(error => {
                        console.error('Failed to send data:', error);
                    });
            };
            reader.readAsDataURL(avatarFile);
        } else {
            const data = {
                nick_name: newName,
                description: newDescription,
                avatar: ''
            };

            fetch(`/changeProfile/${userId}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            })
                .then(response => {
                    if (!response.ok) {
                        throw new Error('Failed to update profile');
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
                })
                .catch(error => {
                    console.error('Failed to send data:', error);
                });
        }

        console.log('New Name:', newName);
        console.log('New Description:', newDescription);
    });
});