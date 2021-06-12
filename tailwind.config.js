module.exports = {
  future: {
    removeDeprecatedGapUtilities: true,
    purgeLayersByDefault: true,
  },
  purge: {
    mode: 'all',
    content: [
      './ui/**/*.jsx',
      './ui/**/*.html',
    ]
},
  theme: {
    extend: {
      minHeight: {
        '1/2': '50%',
      },
      width: {
        72: "18rem",
        80: "20rem",
        88: "22rem",
        96: "24rem",
      },
      margin: {
        72: "18rem",
        80: "20rem",
        88: "22rem",
        96: "24rem",
      }
    },
    colors: {
      "gray": {
        dark: "#313030",
        medium: "#403F40",
        light: "#8E8E8E"
      },
      "white": "#F9F9F9",
      "dirty-white": "#ffe6c0",
      "green": "#5EB663",
      "green-light": "#92e897",
      "blue": "#5C8FFF",
      "blue-light": "#709DFF",
      "red": "#FE5A5A",
      "red-light": "#FF9B9B",
      "orange": "#E39827",
      "black": "#1C1C1C"
    },
    boxShadow: {
      default: '0 1px 3px 0 rgba(0, 0, 0, .1), 0 1px 2px 0 rgba(0, 0, 0, .06)',
      md: '0 4px 6px -1px rgba(0, 0, 0, .1), 0 2px 4px -1px rgba(0, 0, 0, .06)',
      lg: '0 10px 15px -3px rgba(0, 0, 0, .1), 0 4px 6px -2px rgba(0, 0, 0, .05)',
      xl: '0 20px 25px -5px rgba(0, 0, 0, .1), 0 10px 10px -5px rgba(0, 0, 0, .04)',
      "2xl": '0 25px 50px -12px rgba(0, 0, 0, .25)',
      "3xl": '0 35px 60px -15px rgba(0, 0, 0, .3)',
      inner: 'inset 0 4px 8px 0 rgba(0, 0, 0, 0.9)',
      outline: '0 0 0 3px rgba(66, 153, 225, 0.5)',
      'none': 'none',
}
  },
  variants: {},
  plugins: [],
}
