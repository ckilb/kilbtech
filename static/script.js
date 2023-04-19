AOS.init();

const contactForm = document.querySelector('#contact form');

contactForm.addEventListener('submit', (event) => {
    event.preventDefault();

    const data = new FormData(contactForm);

    fetch(contactForm.getAttribute('action'), {
        method: contactForm.getAttribute('method'),
        body: JSON.stringify({
            email: data.get('email'),
            message: data.get('message')
        })
    }).then(res=> {
       alert(res.status === 204);
    });
})
