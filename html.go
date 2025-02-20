<!DOCTYPE html>
<html lang="ar" dir="rtl">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>𝐋𝐎𝐕𝐈𝐍𝐆 𝐅𝐀𝐌𝐈𝐋𝐘 - اقتراحات الديسكورد</title>
  <style>
    body {
      font-family: Arial, sans-serif;
      background-color: #f4f4f9;
      margin: 0;
      padding: 0;
      display: flex;
      flex-direction: column;
      align-items: center;
      padding: 20px;
    }

    header {
      text-align: center;
      margin-bottom: 20px;
    }

    h1 {
      color: #4a4a8a;
      font-size: 28px;
      margin-bottom: 5px;
    }

    a.discord-link {
      background-color: #5865F2;
      color: white;
      padding: 8px 16px;
      border-radius: 8px;
      text-decoration: none;
      font-weight: bold;
      transition: background-color 0.3s;
    }

    a.discord-link:hover {
      background-color: #404ecb;
    }

    form {
      background: #fff;
      padding: 20px;
      border-radius: 12px;
      box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
      width: 100%;
      max-width: 500px;
      margin-bottom: 20px;
    }

    textarea {
      width: 100%;
      padding: 10px;
      border: 1px solid #ccc;
      border-radius: 8px;
      resize: none;
      height: 100px;
      margin-bottom: 10px;
    }

    button {
      background-color: #4a90e2;
      color: white;
      padding: 10px 20px;
      border: none;
      border-radius: 8px;
      cursor: pointer;
      transition: background-color 0.3s, transform 0.2s;
    }

    button:hover {
      background-color: #357ab8;
      transform: scale(1.05);
    }

    .message {
      margin-top: 10px;
      padding: 10px;
      background-color: #d4edda;
      color: #155724;
      border-radius: 8px;
      font-size: 14px;
      display: none;
      text-align: center;
      width: 100%;
      max-width: 500px;
    }

    .suggestions {
      width: 100%;
      max-width: 500px;
      list-style: none;
      padding: 0;
    }

    .suggestion {
      background: #fff;
      padding: 15px;
      margin-bottom: 10px;
      border-radius: 12px;
      box-shadow: 0 3px 5px rgba(0, 0, 0, 0.1);
      position: relative;
    }

    .timestamp {
      font-size: 12px;
      color: #888;
      position: absolute;
      bottom: 8px;
      left: 10px;
    }
  </style>
</head>
<body>
  <header>
    <h1>𝐋𝐎𝐕𝐈𝐍𝐆 𝐅𝐀𝐌𝐈𝐋𝐘</h1>
    <a class="discord-link" href="https://discord.gg/HYAEu3zK" target="_blank">انضم إلى الديسكورد 🎮</a>
  </header>

  <h2>📢 أرسل اقتراحك لسيرفر الديسكورد</h2>
  <form id="suggestionForm">
    <textarea id="suggestionInput" placeholder="اكتب اقتراحك هنا..." required></textarea>
    <button type="submit">إرسال الاقتراح</button>
  </form>

  <div class="message" id="successMessage">✅ تم حفظ اقتراحك بنجاح وسيظل محفوظًا دائمًا!</div>

  <h2>💡 الاقتراحات:</h2>
  <ul id="suggestionsList" class="suggestions"></ul>

  <script>
    const form = document.getElementById('suggestionForm');
    const input = document.getElementById('suggestionInput');
    const suggestionsList = document.getElementById('suggestionsList');
    const successMessage = document.getElementById('successMessage');

    // تحميل الاقتراحات المحفوظة من التخزين المحلي إلى الأبد عند تحميل الصفحة
    window.onload = function () {
      const savedSuggestions = JSON.parse(localStorage.getItem('suggestionsForever')) || [];
      savedSuggestions.forEach(addSuggestionToList);
    };

    form.addEventListener('submit', function (e) {
      e.preventDefault();
      const suggestionText = input.value.trim();
      if (suggestionText) {
        const timestamp = new Date().toLocaleString('ar-EG');
        const suggestion = { text: suggestionText, date: timestamp };

        addSuggestionToList(suggestion);
        saveSuggestionForever(suggestion);
        showSuccessMessage();
        input.value = '';
      }
    });

    function addSuggestionToList(suggestion) {
      const li = document.createElement('li');
      li.className = 'suggestion';
      li.innerHTML = `
        <p>${suggestion.text}</p>
        <span class="timestamp">📅 ${suggestion.date}</span>
      `;
      suggestionsList.prepend(li); // إضافة الاقتراح في الأعلى
    }

    function saveSuggestionForever(suggestion) {
      const suggestions = JSON.parse(localStorage.getItem('suggestionsForever')) || [];
      suggestions.unshift(suggestion); // إضافة الاقتراح مع التاريخ
      localStorage.setItem('suggestionsForever', JSON.stringify(suggestions));
    }

    function showSuccessMessage() {
      successMessage.style.display = 'block';
      setTimeout(() => {
        successMessage.style.display = 'none';
      }, 3000); // إخفاء الرسالة بعد 3 ثوانٍ
    }
  </script>
</body>
</html>
