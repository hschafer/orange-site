export function readableTimestamp(createdOn) {
    function message(amount, unit) {
        amount = Math.floor(amount)
        return amount + " " + unit + (amount == 1 ? "" : "s") + " ago"
    }

    // Given the timestamp a post was created, returns a human-readable timestamp of the form
    // "6 seconds ago", "7 minutes ago", "4 hours ago", "2 days ago"
    const timestamp = new Date(createdOn);
    const now = new Date();
    const millis = now - timestamp;   // In ms

    if (millis < 1000) {  // Less than a second
        return "now";
    }

    const seconds = millis / 1000;
    if (seconds < 60) {
        return message(seconds, "second")
    }

    const minutes = seconds / 60;
    if (minutes < 60) {
        return message(minutes, "minute")
    }

    const hours = minutes / 60;
    if (hours < 24) {
        return message(hours, "hour")
    }

    const days = hours / 24;
    return message(days, "day")
}