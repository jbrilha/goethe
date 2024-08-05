/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [ "./**/*.html", "./**/*.templ", "./**/*.go", ],
	safelist: [],
    theme: {
        extend: {
            fontFamily: {
                appleChancery: ['AppleChancery', 'sans-serif'],
                academyEngraved: ['AcademyEngraved', 'sans-serif'],
            },
        },
    },
}
