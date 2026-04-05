(function() {
    'use strict';

    if (!window.location.pathname.startsWith("/usedcar")) {
        console.log("AdCar: processAutocomCars — not on usedcar page, skipping");
        return;
    }    

    function toSnakeCase(text) {
        return text.replace(/\s+/g, ' ')
            .trim()
            .replace(/([a-z])([A-Z])/g, '$1_$2')
            .replace(/[\s\-]+/g, '_')
            .toLowerCase();
    }

    function getText(selector) {
        const element = document.querySelector(selector);
        return element ? element.textContent.trim() : '';
    }

    async function processAutocomCars() {
        console.log("AdCar: processAutocomCars started");

        // Загружаем настройки
        const result = await chrome.storage.sync.get(['department', 'adcarHost']);
        const settings = {
            department: result.department || 'Импорт Япония',
            adcarHost: result.adcarHost || 'http://localhost:8080/api/v1/cars/'
        };

        console.log('🔧 Settings:', settings); // Дебаг

        const data = {
            source: "autocom",
            department: settings.department,
            link: window.location.href,
            steering_side: "Right",
            photos: []
        };

        // Фото
        const photos = document.querySelectorAll('#photo #slider-photo img, .photo img');
        photos.forEach(img => {
            const src = img.src || img.dataset.src;
            if (src && src.includes('autocj.co.jp/photo')) {
                data.photos.push(src);
            }
        });

        // Характеристики (улучшенный парсер)
        const features = document.querySelectorAll('.vinfo dt, .vinfo dd');
        for (let i = 0; i < features.length; i++) {
            const label = features[i].innerText.trim();
            const value = features[i + 1]?.innerText?.trim() || '';

            if (['Производитель', 'Make', 'Произв'].includes(label)) data.make = value;
            if (['Модель', 'Name'].includes(label)) data.model = value;
            if (['Год / месяц R', 'R Year / Month'].includes(label)) {
                const prodDate = value.split(' / ');
                data.year = parseInt(prodDate[0] || 0);
                data.month = parseInt(prodDate[1] || 0);
            }
            if (['Пробег', 'Mileage'].includes(label)) {
                data.mileage = parseInt(value.replace(/[,\\s]/g, '')) || 0;
            }
            if (['Объем', 'CC'].includes(label)) {
                data.engine_capacity = parseInt(value) || 0;
            }
            if (['Топливо', 'Fuel'].includes(label)) data.fuel = value;
            if (['Передача', 'Transmission'].includes(label)) data.transmission = value;
            if (['Привод', 'Drive'].includes(label)) data.drive = value;
            if (['Цвет', 'Color'].includes(label)) data.color = value;
            if (label === 'Grade') data.trim = value;
            if (['Тип кузова', 'Body Type'].includes(label)) data.body_type = value;

            // Horsepower
            if (label.includes('PS') || label.includes('HP') || label.includes('kW')) {
                data.horse_power = parseInt(value) || 0;
            }
        }

        // Цена
        const prices = document.querySelectorAll('.pricelist dd');
        const priceText = prices[prices.length - 1]?.innerText?.replaceAll(/[,\\s]/g, '') || '';
        data.currency = priceText.startsWith('$') ? 'USD' : 'JPY';
        data.price = parseInt(priceText.replace(/[^\d]/g, '')) || 0;

        // AdCar объект
        const adcarCar = {
            car_id: Date.now(),
            department: data.department,
            resource: 'autocj.co.jp',
            from_country: 'JP',
            link: data.link,
            vehicle_type: 'passenger',
            vin: '',
            make: data.make,
            model: data.model,
            month: data.month,
            year: data.year,
            age: new Date().getFullYear() - data.year,
            body_type: data.body_type,
            is_right_steering: true,
            color: data.color,
            trim: data.trim,
            mileage: data.mileage,
            fuel: data.fuel,
            engine_volume: data.engine_capacity,
            horse_power: data.horse_power || 0,
            transmission: data.transmission,
            drive_type: data.drive,
            photos: data.photos.slice(0, 10),
            price: data.price,
            price_currency: data.currency,
            additional_context: JSON.stringify(data)
        };

        console.log('AdCar данные:', adcarCar);

        // Кнопка
        if (document.querySelector('.adcar-save')) return;

        const sendButton = document.createElement('div');
        sendButton.className = 'adcar-save';

        // Стиль как у daecha: фиксированная, маленькая, градиент, без тяжёлого blur
        sendButton.style.cssText = `
            position: fixed;
            top: 5px;
            left: 5px;
            z-index: 2147483647;

            width: 140px;
            height: 40px;

            background: linear-gradient(135deg, #007bff, #0056b3);
            color: white;
            font-weight: 600;
            font-size: 13px;

            border-radius: 12px;
            border: none;
            outline: none;
            cursor: pointer;

            display: flex;
            align-items: center;
            justify-content: center;
            white-space: pre-line;
            text-align: center;
            line-height: 1.2;

            transition: all 0.3s ease;
            box-shadow: 0 8px 25px rgba(0, 123, 255, 0.4);
        `;

        // Анимация при наведении
        sendButton.addEventListener('mouseenter', () => {
            sendButton.style.transform = 'translateY(-2px)';
            sendButton.style.boxShadow = '0 12px 30px rgba(0, 123, 255, 0.5)';
        });
        sendButton.addEventListener('mouseleave', () => {
            sendButton.style.transform = '';
            sendButton.style.boxShadow = '0 8px 25px rgba(0, 123, 255, 0.4)';
        });

        // Текст кнопки — можно оставить с иконкой, если хочешь, или просто текст
        sendButton.innerText = 'Сохранить\nв AdCar-CRM';

        // Логика отправки — остаётся почти та же, меняем только визуал
        sendButton.onclick = () => {
            sendButton.style.background = 'linear-gradient(135deg, #2196F3, #1976D2)';
            sendButton.style.transform = 'none';            

            chrome.runtime.sendMessage(
                {
                    action: 'saveCar',
                    data: adcarCar,
                    adcarHost: settings.adcarHost,
                },
                (response) => {
                    if (chrome.runtime.lastError) {
                        console.error('Extension error:', chrome.runtime.lastError.message);
                        sendButton.style.background = 'linear-gradient(135deg, #f44336, #d32f2f)';
                        sendButton.innerText = 'Ошибка\nрасширения';
                    } else if (response?.action === 'carSaved') {
                        console.log('Сохранено!', response.result);
                        sendButton.style.background = 'linear-gradient(135deg, #28a745, #1e7e34)';
                        sendButton.innerText = 'Сохранено!';
                        setTimeout(() => sendButton.remove(), 2000);
                    } else if (response?.action === 'saveError') {
                        console.error('Server error:', response.error);
                        sendButton.style.background = 'linear-gradient(135deg, #f44336, #d32f2f)';
                        sendButton.innerText = 'Ошибка\nсервера';
                    } else {
                        sendButton.style.background = 'linear-gradient(135deg, #f44336, #d32f2f)';
                        sendButton.innerText = 'Неизвестная\nошибка';
                    }

                    // Сброс через 3 секунды при ошибке
                    setTimeout(() => {
                        if (
                            sendButton.parentNode &&
                            (
                                sendButton.innerText.includes('Ошибка') ||
                                sendButton.innerText.includes('Неизвестная')
                            )
                        ) {
                            sendButton.style.background = 'linear-gradient(135deg, #007bff, #0056b3)';
                            sendButton.style.transform = '';
                            sendButton.innerText = 'Сохранить\nв AdCar';
                        }
                    }, 3000);
                }
            );
        };

        document.body.appendChild(sendButton);
    }

    // Автозапуск
    if (document.readyState === 'loading') {
        document.addEventListener('DOMContentLoaded', () => setTimeout(processAutocomCars, 1000));
    } else {
        setTimeout(processAutocomCars, 1000);
    }
})();