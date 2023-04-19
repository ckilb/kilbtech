AOS.init();

const contactForm = document.querySelector('#contact form');
const contactFormButton = contactForm.querySelector('button');
const contactFormButtonSpinner = contactFormButton.querySelector('.animate-spin');

contactForm.addEventListener('submit', (event) => {
    event.preventDefault();

    const data = new FormData(contactForm);

    contactForm.classList.add('pointer-events-none');
    contactForm.classList.add('opacity-75');
    contactFormButton.classList.add('text-[0px]');
    contactFormButtonSpinner.classList.remove('hidden');
    contactFormButtonSpinner.classList.add('inline-block');

    fetch(contactForm.getAttribute('action'), {
        method: contactForm.getAttribute('method'),
        body: JSON.stringify({
            email: data.get('email'),
            message: data.get('message')
        })
    }).then(res=> {

    }).finally(() => {
        contactForm.classList.remove('pointer-events-none');
        contactForm.classList.remove('opacity-75');
        contactFormButtonSpinner.classList.remove('inline-block');
        contactFormButtonSpinner.classList.add('hidden');
    });
})
