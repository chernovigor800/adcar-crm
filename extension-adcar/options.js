document.addEventListener('DOMContentLoaded', loadSettings);
document.getElementById('save').onclick = saveSettings;

async function saveSettings() {
    const department = document.getElementById('department').value.trim();
    const adcarHost = document.getElementById('host').value.trim();
    if (!department || !adcarHost) { showStatus('❌ Заполните все поля', '#f44336'); return; }
    try {
        await chrome.storage.sync.set({department, adcarHost});
        showStatus('✅ Сохранено!', '#4caf50');
    } catch(err) { showStatus('❌ Ошибка', '#f44336'); }
}

async function loadSettings() {
    try {
        const result = await chrome.storage.sync.get(['department','adcarHost']);
        document.getElementById('department').value = result.department || 'NSK';
        document.getElementById('host').value = result.adcarHost || 'http://localhost:8080/api/v1/cars/';
    } catch(err) { console.error(err); }
}

function showStatus(msg, color) {
    const status = document.getElementById('status');
    status.textContent = msg; status.style.color = color;
    setTimeout(() => status.textContent = '', 3000);
}