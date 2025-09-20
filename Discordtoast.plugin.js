module.exports = class PeriodicToast {
  // The timer variable will store the interval ID so we can stop it later.
  timer;

  let y = BdApi.Webpack.getStore("PresenceStore").getStatus('450867169581072394')
  start() {
    this.timer = setInterval(() => {
      x = BdApi.Webpack.getStore("PresenceStore").getStatus('450867169581072394')
      if (x != y)
        BdApi.UI.showToast(BdApi.Webpack.getStore("PresenceStore").getStatus('450867169581072394'), {
          type: "info",
          timeout: 5000
          y = x
      });
    }, 10000); // 10000 milliseconds = 10 seconds
    
    // Optional: show a toast when the plugin starts to confirm it's working
    BdApi.UI.showToast("PeriodicToast plugin started!", { type: "success" });
  }

  stop() {
    // Clear the interval to prevent the toasts from continuing after the plugin is disabled.
    clearInterval(this.timer);
    BdApi.UI.showToast("PeriodicToast plugin stopped.", { type: "error" });
  }
};
