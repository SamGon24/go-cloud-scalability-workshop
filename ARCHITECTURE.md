# 0:55 — Architecture Design Activity

**Time:** 15 minutes (10 design + 5 share)

---

## Scenario (read to students)

> "Your app just went viral. Users upload images from their phones. Each image needs to be resized to 3 thumbnails and tagged by an ML model that takes ~5 seconds.
>
> You currently have **one server** that does everything inline. Twitter just discovered you. Ten thousand uploads land in the next hour.
>
> Design a system that won't fall over."

---

## Teams must answer (put these on the whiteboard)

1. Where does the upload land? Who handles it?
2. Where does the 5-second ML work happen? **Not** in the HTTP handler, right?
3. What component decouples fast uploads from slow ML?
4. Where does the image data live? (Hint: not in RAM.)
5. How do you scale each piece independently?
6. What happens if the ML worker crashes mid-job?

---
