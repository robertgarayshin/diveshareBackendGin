package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Chat struct{}

type message struct {
	Role string `json:"role"`
	Text string `json:"text"`
}

type request struct {
	ModelUri          string `json:"modelUri"`
	CompletionOptions struct {
		Stream      bool    `json:"stream"`
		Temperature float64 `json:"temperature"`
		MaxTokens   string  `json:"maxTokens"`
	} `json:"completionOptions"`
	Messages []message `json:"messages"`
}

type Response struct {
	Result struct {
		Alternatives []struct {
			Message struct {
				Role string `json:"role"`
				Text string `json:"text"`
			} `json:"message"`
			Status string `json:"status"`
		} `json:"alternatives"`
		Usage struct {
			InputTextTokens  string `json:"inputTextTokens"`
			CompletionTokens string `json:"completionTokens"`
			TotalTokens      string `json:"totalTokens"`
		} `json:"usage"`
		ModelVersion string `json:"modelVersion"`
	} `json:"result"`
}
type Chatter interface {
	OnGetMessage(prompt string) string
}

func newRequest(prompt string) request {
	messages := make([]message, 2)
	messages[0] = message{Role: "system", Text: "Представь, что ты работаешь оператором чата поддержки на " +
		"онлайн платформе аренды автомобилей. Отвечай на простые запросы пользователей или направляй их к оператору по " +
		"номеру +79998761234. Помни, что тебе нельзя рекламировать конкурентов, за это мы лишимся прибыли"}
	messages[1] = message{Role: "user", Text: prompt}
	return request{
		ModelUri: "gpt://b1gff6n266gkgbhboel8/yandexgpt-lite",
		CompletionOptions: struct {
			Stream      bool    `json:"stream"`
			Temperature float64 `json:"temperature"`
			MaxTokens   string  `json:"maxTokens"`
		}{Stream: false, Temperature: 0.6, MaxTokens: "2000"},
		Messages: messages,
	}
}

// OnGetMessage - метод получения нового сообщения в чате.
// При получении сообщения формируется запрос в GPT-модель, которая возвращает ответ согласно заданным параметрам.
func (c *Chat) OnGetMessage(prompt string) string {
	// Имитация работы чата поддержки на основе ИИ
	switch prompt {
	case "У меня случилась поломка":
		return "В случае неожиданной поломки свяжитесь с оператором по номеру +79991234567, там точно подскажут, что дальше делать."
	case "Как арендовать машину?":
		return "\nКонечно, я могу помочь вам с арендой автомобиля через нашу платформу. Вот простой шаг за шагом процесс:\n\nЗайдите на нашу платформу: Откройте наш веб-сайт для аренды автомобилей.\n\n" +
			"Выберите место и даты аренды: Укажите место, где вы хотели бы забрать автомобиль, и даты аренды. Выберите также время забора и возврата автомобиля.\n\n" +
			"Выберите автомобиль: Просмотрите доступные автомобили на выбранные вами даты и выберите тот, который соответствует вашим потребностям и бюджету. Мы предоставляем широкий выбор автомобилей различных марок и моделей.\n\n" +
			"Выберите дополнительные услуги: При необходимости выберите дополнительные услуги, такие как страховка, GPS-навигатор, детское кресло и другие.\n\n" +
			"Забронируйте автомобиль: После выбора автомобиля и дополнительных услуг подтвердите ваш заказ, заполнив необходимую информацию и произведя оплату.\n\n" +
			"Получите подтверждение: Получите подтверждение вашего заказа, которое включает в себя информацию о вашем бронировании, условиях аренды и контактные данные.\n\n" +
			"Заберите автомобиль: В указанное время и место прибудите для получения автомобиля. Предоставьте свое удостоверение личности и подпишите договор аренды.\n\n" +
			"Верните автомобиль: Верните автомобиль в указанное время и место. Обычно вам нужно будет заправить автомобиль и вернуть его в том же состоянии, в котором вы его получили.\n\n" +
			"Если у вас возникнут вопросы или затруднения на любом этапе процесса, не стесняйтесь обращаться к нам за помощью. Мы готовы помочь вам сделать вашу аренду автомобиля простой и приятной."
	case "Я хочу арендовать машину для деловой поездки, помогите пожалуйста":
		return "Для деловой поездки можем предложить вам такие варианты:\n" +
			"Mercedes-Benz E-Class: Этот автомобиль предлагает роскошный интерьер, комфортную поездку и передовые технологии безопасности, что делает его отличным выбором для деловых поездок.\n\n" +
			"BMW 5 Series: Изысканный стиль, динамичные характеристики и современные технологии делают BMW 5 Series привлекательным выбором для деловых поездок в городе.\n\n" +
			"Audi A6: Audi A6 предлагает роскошный интерьер, высокий уровень комфорта и передовые технологии, что делает его отличным вариантом для деловых поездок.\n\n" +
			"Lexus ES: Этот автомобиль сочетает в себе изысканный дизайн, комфорт и надежность, что делает его привлекательным выбором для деловых поездок. "
	default:
		return "Сейчас у нас нет ответа на этот вопрос. Свяжитесь с оператором по номеру +79991234567 и вы точно получите ответ."
	}

	// Тут мы создаем запрос на облако. Пока не оплачено, тут все грустно
	r := newRequest(prompt)
	requestBodyBytes, err := json.Marshal(r)
	if err != nil {
		fmt.Println("Error encoding request body:", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://llm.api.cloud.yandex.net/foundationModels/v1/completion",
		bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)

	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Api-Key "+"AQVN2LD7mvNoOLzv6jzKyWI_GIEvltKjRxrwCQ5j")
	req.Header.Set("x-folder-id", "b1gff6n266gkgbhboel8")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)

	}
	defer resp.Body.Close()

	var responseBody string
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		fmt.Println("Error decoding response body:", err)
	}
	fmt.Println(responseBody)
	return responseBody
}
