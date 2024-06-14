function updateNavigationLinks() {
    // Получаем текущий URL
    const currentUrl = window.location.href;

    // Используем регулярное выражение для извлечения числа из URL
    const regex = /\/h\/(\d+)/;
    const match = currentUrl.match(regex);

    if (match) {
        // Получаем текущее число страницы
        const currentPage = parseInt(match[1]);

        // Вычисляем номера для страниц "Назад" и "Вперед"
        const prevPage = currentPage > 1 ? currentPage - 1 : 1;
        const nextPage = currentPage + 1;

        // Формируем новые URL
        const prevUrl = currentUrl.replace(regex, `/h/${prevPage}`);
        const nextUrl = currentUrl.replace(regex, `/h/${nextPage}`);

        // Обновляем атрибуты href ссылок
        document.getElementById('backButton').href = prevUrl;
        document.getElementById('forwardButton').href = nextUrl;
    } else {
        console.error('Current URL does not match expected pattern.');
    }
}

// Обновляем ссылки при загрузке страницы
window.onload = updateNavigationLinks;