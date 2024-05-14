document.addEventListener("DOMContentLoaded", function () {
  console.log("Document is loaded!");

  const path = window.location.pathname;

  sendStuff({
    path,
    bucketID: "somethingRandnom",
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
