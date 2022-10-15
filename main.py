import requests
import bs4 as bs


URL = "https://www.serebii.net/pokemon/gen1pokemon.shtml"
resp = requests.get(URL)
soup = bs.BeautifulSoup(resp.text, features="html.parser")
table = soup.find('div', id="content")

pokemons = []
rows = table.findAll('tr')[2:304]
print(len(rows))
for row in rows:
    cell = row.find_all('td')
    try:
        pid = cell[0].get_text().strip()
        image = cell[1].td.a.img['src']
        name = cell[3].a.get_text()
        types = [ img['href'].split('/')[-1] for img in cell[4].find_all('a') ]
        ability = cell[5].get_text().strip()
        base_stats = {}
        base_stats["healthPoints"] = cell[6].get_text()
        base_stats["attack"] = cell[7].get_text()
        base_stats["defense"] = cell[8].get_text()
        base_stats["speedAttack"] = cell[9].get_text()
        base_stats["speedDefense"] = cell[10].get_text()
        base_stats["specialDefense"] = cell[11].get_text()

        pokemons.append({
                        "name": name,
                        "data": {
                            "id": pid,
                            "name": name,
                            "image": image,
                            "types": types,
                            "ability": ability,
                            "baseStats": base_stats}})
    except:
        continue

print(pokemons)
print(len(pokemons))


with open('pokemons.json', 'w') as f:
    f.write(str(pokemons))