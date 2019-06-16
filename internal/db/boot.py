import sqlite3

conn = sqlite3.connect("data.db")

c = conn.cursor()

c.execute("""CREATE TABLE jokes(uid INTEGER PRIMARY KEY AUTOINCREMENT, content TEXT)""")

with open("jokes.txt", "rt") as f:
    for l in f:
        l = l.replace("\\n", "\n")
        c.execute("INSERT INTO jokes(content) VALUES (?)", (l,))

conn.commit()
conn.close()
