# أداة إرسال HTML كبريد إلكتروني (GoMailer)

## نظرة عامة

هي أداة بسيطة كتبتها بلغة Go لإرسال رسائل بريد إلكتروني باستخدام SMTP. تدعم الأداة إرسال ملفات HTML وتتيح تخصيص معلومات الاتصال بخادم SMTP من خلال الـ environment variables.

## المميزات

- إرسال بريد إلكتروني باستخدام بروتوكول SMTP مع دعم TLS/SSL
- إرسال ملفات HTML
- تكوين سهل باستخدام متغيرات البيئة
- دعم مُختلَف الـ SMTP clients (الإعدادات الافتراضية تناسب Gmail)

## متطلبات التشغيل

- يمكنك استخدام Docker ببساطة، لست بحاجة لتثبيت أي شيء على جهازك (باستثناء Docker طبعًا 😄)
- إذا أردت تشغيل التطبيق على جهازك يجب أن تُحمّل Go 1.24
- سوف تحتاج أيضًا لبريد إلكتروني مع الصلاحيات اللازمة لاستخدام SMTP (بريدك الشخصي على Gmail يفي بالغرض)

## هيكل المشروع

```
gomailer/
├── cmd/
│   └── gomailer/
│       └── main.go        # application entry point
├── internal/
│   ├── config/            # configuration handling
│   ├── email/             # email sending functionality
│   └── templates/         # template loading functionality
├── templates/
│   └── email.html         # HTML email template
├── .dockerignore
├── Dockerfile
├── .env.example           # blueprint of what your .env should include
├── go.mod
├── Dockerfile
└── README.md
```

## تحميل الـ repo على جهازك

```bash
git clone https://github.com/ahmezing/gomailer.git
cd gomailer
```

## Environment Variables

قم بإنشاء ملف `.env` وعيّن القيم التالية:

```
EMAIL=your-email@gmail.com
EMAIL_PASSWORD=your-password-or-app-password
TO_EMAIL=recipient@example.com
SMTP_HOST=smtp.gmail.com
SMTP_PORT=465
```

## التشغيل باستخدام Docker

يمكنك بناء وتشغيل التطبيق باستخدام Docker:

```bash
# build the image
docker build -t gomailer .
# run the container with the environment variables
docker run --env-file .env gomailer
```

## التشغيل على جهازك

```bash
# Export environment variables from .env
export $(cat .env | xargs)
# Run the application
go run cmd/gomailer/main.go
```

## إرسال رسائل بريد إلكتروني مخصصة

- قم بتعديل قالب HTML في `templates/email.html`

## ملاحظات هامة إذا كنت تستخدم Gmail

تحتاج إلى:

1. تمكين وصول التطبيقات الأقل أمانًا في إعدادات حسابك
2. أو استخدام كلمة مرور التطبيق «App Password» إذا كنت تستخدم المصادقة الثنائية «Two Factor Authentication»

## استكشاف الأخطاء وإصلاحها

### المشاكل الشائعة:

1. **Authentication Failed**:

تأكد من استخدام البريد الإلكتروني وكلمة المرور الصحيحين. بالنسبة لـ Gmail، استخدم كلمة مرور التطبيق.

2. **Connection Refused**:

تحقق من صحة خادم SMTP والـport.

3. **SSL/TLS Error**:

تأكد من أن port 465 غير محظور بواسطة جدار الحماية الخاص بك.
