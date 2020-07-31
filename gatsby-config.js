module.exports = {
  siteMetadata: {
    title: "Bro Long",
  },
  plugins: [
    `gatsby-plugin-react-helmet`,
    {
      resolve: `gatsby-plugin-manifest`,
      options: {
        name: "Bro Long",
        short_name: "Bro Long",
        start_url: "/",
        icon: "src/images/star.png",
      },
    },
  ],
};
