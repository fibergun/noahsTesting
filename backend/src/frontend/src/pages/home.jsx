import { useSession } from "../routers/Session.jsx";

function capitalize(str) {
  if (!str) return str; // handle empty/undefined safely
  return str.charAt(0).toUpperCase() + str.slice(1);
}

function Home() {
  const { session } = useSession();

  return (
    <div>
      <h1>Welcome {capitalize(session.username)}</h1>
      <h3>Uitleg</h3>
      <p>
        Tijdens dit weekend probeer je zo veel mogelijk punten te verzamelen,
        deze punten kan je alleen maar verdien door taken uit te voeren. Taken
        zijn kleine opdrachten die door jou of door andere mensen die mee zijn
        op dit weekend zijn bedacht. <br/> Onderaan deze pagina staat een korte uitleg
        van hoe deze app verder werkt, mocht dit niet meteen duidelijk zijn! Je
        mag zo veel taken actief hebben als je maar wilt, maar weet wel dat niet
        afgemaakte taken je punten kosten in plaats van geven! <br/> Degene die aan
        het einde van het weekend de meeste punten heeft verdient hiermee naast
        de eeuwige roem, ook nog een mooie prijs!
      </p>
      <h3>Regels</h3>
      <p>
        Probeer taken die je maakt relatief kort te houden, iets wat iemand
        meteen, in 15 minuten of in 30 minuten uit kan voeren, ook iets wat
        iemand verzameld/voorbereid moet hebben aan het einde van het weekend
        werkt goed.</p> <p>  Probeer er voor te zorgen dat alle taken daadwerkelijk
        mogelijk zijn over de (resterende) duur van het weekend! </p><p> Taken zijn
        niet uniek in de groep, dus 2 mensen kunnen dezelfde taak krijgen,
        probeer te zorgen dat de taken door meerdere mensen uitvoerbaar zijn,
        tenzij dit direct tegen het idee van de taak in gaat natuurlijk. </p><p> Probeer
        er aan te denken dat een taak door iedereen uitvoerbaar moet zijn.
      </p>
      <h3>Uitleg</h3>
      <p>
        De ping pagina doet niets, dit is alleen om te kijken of de backend nog
        draait, mocht dit ooit een error geven, begin met bidden. </p><p> De "make task"
        pagina laat je een taak aanmaken, klik op de pagina, voer een taak in en
        klik op "submit task". Je mag zo veel regels maken als je maar wilt,
        maar denk er aan dat je ze zelf ook kunt krijgen! </p><p> De "Get Random Task"
        pagina geeft je, als er een beschikbaar is direct een nieuwe regel, deze
        word weergeven en zal toegevoegd worden aan je takenlijst. </p><p> De "All
        Tasks" pagina laat je punten zien, en alle taken die je hebt gedaan en
        open hebt staan. </p><p> Je krijgt een minpunt wanneer je een taak open hebt
        staan, wat een pluspunt word wanneer je hem gedaan hebt, dus wees niet
        bang als je punten in de min staan, dit kan snel omslaan als je een paar
        taken voltooid. <br/> Klik simpelweg op de "complete" knop om de taak te
        voltooien..
      </p>
    </div>
  );
}

export default Home;
