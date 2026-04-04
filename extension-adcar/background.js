chrome.runtime.onMessage.addListener((request, sender, sendResponse) => {
  if (request.action === 'saveCar') {
    const host = request.adcarHost || 'http://localhost:8080/';
    
    (async () => {
      try {
        console.log('Отправка на:', host);
        const response = await fetch(host, {
          method: 'POST',
          headers: { 
            'Content-Type': 'application/json',
            'Accept': 'application/json'
          },
          body: JSON.stringify(request.data)
        });
        
        if (!response.ok) throw new Error(`HTTP ${response.status}`);
        const data = await response.json();
        
        sendResponse({ action: 'carSaved', result: data });
      } catch (error) {
        console.error('❌ Background error:', error);
        sendResponse({ action: 'saveError', error: error.message });
      }
    })();
    
    return true; // Обязательно для async ответа!
  }
});