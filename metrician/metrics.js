document.addEventListener("DOMContentLoaded", function () {
  console.log("Document is loaded!");
  // so now we should figure what kind of data
  // we want to send to the server

  // maybe path ip bucketID anything else

  const path = window.location.pathname;

  sendStuff({
    path,
    bucketID: "somethingradnom",
  });
});

async function sendStuff(stuf) {
  var url = "http://localhost:30001/data";

  const result = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(stuf),
  });

  console.log(result);
}
