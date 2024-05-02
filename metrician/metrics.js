document.addEventListener("DOMContentLoaded", function () {
  console.log("Document is loaded!");
  sendStuff({});
});

async function sendStuff(stuf) {
  var url = "http://localhost:30001/data";
  var params = { foo: "bar" };

  const result = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(params),
  });

  console.log(result);
}
