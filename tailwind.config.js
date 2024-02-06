/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ['./view/**/*.templ'],
	theme: {
		extend: {
			colors: {
				'background': '#000000',
				'panel': '#121212',
				'feature': '#232323',
				'unfocused': '#a7a7a7',
				'focused': '#ffffff',
				'accent': '#1ed75f',
				'accent-hover': '#11a645',
			},
		},
	},
	plugins: [],
};
