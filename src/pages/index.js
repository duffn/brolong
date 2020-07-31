import React from "react";
import { Helmet } from "react-helmet";
import "../styles/index.css";

function Index() {
  return (
    <main>
      <Helmet>
        <title>Bro Long</title>
      </Helmet>
      <h1>Bro Long.</h1>
      <p className="tag-line">Say hi or bye to bros.</p>
      <p className="pre-block">
        <pre className="pre-command">curl https://brolong/api/</pre>
        <pre>{JSON.stringify({ message: "Hi, broast beef." })}</pre>
      </p>
      <a href="https://nicholasduffy.com">Me.</a>
      <a href="https://github.com/duffn/brolong">GitHub.</a>
    </main>
  );
}

export default Index;
