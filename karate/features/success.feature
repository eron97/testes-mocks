Feature: Listar reservas
    Background: 
        * call read('../utils/BookingAuth.feature')
        * url 'https://treinamento-api.herokuapp.com'
        * header Accept = 'application/json'


    @contract
    Scenario: Garantir o contrato do retorno da lista de reservas
        Given path 'booking'
        When method get
        Then status 200
        And match each response == { bookingid: "#number" }

