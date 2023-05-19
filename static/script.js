AOS.init();

const contactForm = document.querySelector('#contact form');

if (contactForm) {
    const contactFormButton = contactForm.querySelector('button');
    const contactFormButtonSpinner = contactFormButton.querySelector('.animate-spin');
    const contactFormButtonText = contactFormButton.querySelector('.text');
    const contactFormSuccessMessage = contactForm.querySelector('.success-message');
    const contactFormErrorMessage = contactForm.querySelector('.error-message');

    contactForm.addEventListener('submit', (event) => {
        event.preventDefault();

        const data = new FormData(contactForm);

        contactForm.classList.add('pointer-events-none');
        contactForm.classList.add('opacity-75');
        contactFormButton.classList.add('text-[0px]');
        contactFormButtonSpinner.classList.remove('hidden');
        contactFormButtonSpinner.classList.add('inline-block');
        contactFormButtonText.classList.remove('inline-block');
        contactFormButtonText.classList.add('hidden');

        fetch(contactForm.getAttribute('action'), {
            method: contactForm.getAttribute('method'),
            body: JSON.stringify({
                email: data.get('email'),
                message: data.get('message')
            })
        }).then(res=> {
            if (res.status !== 204) {
                contactFormErrorMessage.classList.remove('hidden');

                return;
            }

            contactFormSuccessMessage.classList.remove('hidden');
        }).catch(() => {
            contactFormErrorMessage.classList.remove('hidden');
        }).finally(() => {
            contactFormButton.classList.add('hidden');
        });
    })
}
