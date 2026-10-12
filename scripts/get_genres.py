import sys
import requests
from dotenv import load_dotenv
import os


def main():
    url = "https://api.themoviedb.org/3/genre/movie/list?language=en"

    api_key = sys.argv[1] if len(sys.argv) > 1 else None
    if api_key is None:
        load_dotenv()
        api_key = os.environ.get("TMDB_KEY")
        if api_key is None:
            raise Exception("no api key given")

    headers = {"Authorization": f"Bearer {api_key}"}
    r = requests.get(url, headers=headers)
    r.raise_for_status()

    genres = r.json().get("genres")
    with open("genre_list.csv", "w") as f:
        f.write("name\n")
        for genre in genres:
            f.write(genre.get("name") + "\n")

    return


if __name__ == "__main__":
    main()
