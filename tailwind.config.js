/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [ "./**/*.html", "./**/*.templ", "./**/*.go", ],
	safelist: [],
    theme: {
        extend: {
            fontFamily: {
                appleChancery: ['AppleChancery', 'sans-serif'],
                academyEngraved: ['AcademyEngraved', 'sans-serif'],
                baskerville: ['Baskerville', 'sans-serif'],
                timesNewRoman: ['TimesNewRoman', 'sans-serif'],
                w95: ['W95FA', 'sans-serif'],
            },
        },
    },
}
