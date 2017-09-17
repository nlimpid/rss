package handler

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/nlimpid/rss/models"
)

func Test_getPost(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want models.ZhihuPost
	}{
		// TODO: Add test cases.
		{
			"test1", args{"oh-hard"},
			models.ZhihuPost{
				Name:        "硬派健身",
				Description: "每日一篇质量长文，微信公众：硬派健身。",
				Link:        "/oh-hard",
				Avatar: models.ZhihuPostAvatar{
					ID:       "d8752eeb8ddb0ca382afc836de6224c0",
					Template: "https://pic1.zhimg.com/{id}_{size}.jpg",
				}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getPost(tt.args.name)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPost() = %v, want %v", got, tt.want)
			}
			if err != nil {
				t.Errorf("getPost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getItem(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want models.ZhihuPost
	}{
		{
			"test1", args{"oh-hard"},
			models.ZhihuPost{
				Name:        "硬派健身",
				Description: "每日一篇质量长文，微信公众：硬派健身。",
				Avatar: models.ZhihuPostAvatar{
					ID:       "0a47432f4ef552ceaf1da5a9fe11a443",
					Template: "https://pic4.zhimg.com/{id}_{size}.jpg",
				}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getItem(tt.args.name)
			if !reflect.DeepEqual(got, tt.want) {
				// t.Errorf("getPost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplace(t *testing.T) {
	regex, _ := regexp.Compile(`<img src="(v2-[0-9a-zA-Z]{32}.(png|jpg|gif))"`)
	beforeStr := "</p><img src=\"v2-47487ea266bdf33aa0c0552331ecc3e0.png\" data-rawwidth=\"773\" data-rawheight=\"457\"><p>"
	res := regex.FindAllString(beforeStr, 3)
	fmt.Printf("res = %v\n", res)
	refControl := "https://rss.nlimpid.com/zhihu_image?image="
	replacedstr := regex.ReplaceAllString(beforeStr, fmt.Sprintf("<img src=\"%vhttps://pic4.zhimg.com/$1\"", refControl))
	fmt.Printf("after str = %v\n", replacedstr)
}

func Test_getItems(t *testing.T) {
	type args struct {
		name string
	}
	test1 := models.ZhihuItem{
		Title:       "黑色星期一启示录",
		TitleImage:  "https://pic4.zhimg.com/fab5bc83a276ba7d5f16e951454408f7_r.jpg",
		Description: `1987年10月19日，黑色星期一，美国股市在没有重大利空的情况下大跌，道琼斯指数当天暴跌22%。投资者发现市场突然崩溃的几率并不如他们所想的那么低。也就有了我们所熟知的肥尾现象(Fat-tailed Distribution)。<p><img src="44960e386e5207dc05fc5f7f8f000824" data-rawwidth="585" data-rawheight="403">人们也借此时机认识到对冲下方风险（Downside Risk）的重要性，所以至此以后价外看跌期权的需求量在一般情况下要比价外看涨期权更大，而价外看跌期权倾向于over-priced，因此在实际观察中价外看跌期权的隐含波动率（Implied Volatility）要比相同参数的价外看涨期权高，这就波动率偏移（Volatility Skew）。</p><img src="62e177252328753beaf33b0acc2ad5f2" data-rawwidth="749" data-rawheight="373"><br><p>波动率偏移在大部分指数和普通有期权交易的股票都可以被观察到，偏移方向与程度可以通过风险反转策略(Risk Reversal)来测量。</p><img src="77aa44b23a00ce6bb3cb260c9b95def9" data-rawwidth="545" data-rawheight="230"><p><br>也可以通过观察90%-110% Moneyness期权波动性之间的Spread来衡量</p><p>$SPX - spread为正，多单持有者偏向于买入价外看跌期权对冲下方风险</p><p><img src="0784813baef0d46e9f05d9c006439270" data-rawwidth="736" data-rawheight="527">$FXI - Spread为负，空单持有者偏向于买入价外看涨期权对冲上方风险</p><img src="2f1e3728a6e6a937a67cfed23adfd3b4" data-rawwidth="736" data-rawheight="527"><p><br>1987年黑色星期一的另一个“副产品”是促使投资者开始重视金融产品波动性(Volatility)。CBOE在1993年正式推出建立在SP100之上的波动率指数(VXO)，专门跟踪30天内标普100的隐含波动率；后来演化而成的VIX更是在2008年金融危机中为大众所熟知。</p><p>然而隐含波动率(Implied Volatility)是由衍生品价格推演出来，与由标的物历史价格变化所计算出来的的实际波动率(Realized Volatility)并无直接关系。两者之间所产生的spread称之为Volatility Risk Premium (VRP)。 </p><p><img src="aa149972cd85488512703bb560ff9820" data-rawwidth="736" data-rawheight="527">VRP的出现也反映了投资者愿意付出更高的价格对当前的投资组合进行对冲操作。</p><p>(封面图片Credit:<a href="http://jwintle.com/">jwintle.com </a>2012)</p>`,
		Link:        "/p/20063092",
		Created:     "2015-06-09T12:37:01+08:00"}
	testItem := []models.ZhihuItem{test1}
	tests := []struct {
		name      string
		args      args
		wantItems []models.ZhihuItem
	}{
		// TODO: Add test cases.
		{"test1", args{"option"}, testItem},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotItems := getItems(tt.args.name, "1", "0"); !reflect.DeepEqual(gotItems, tt.wantItems) {
				t.Errorf("getItems() = %v, want %v", gotItems, tt.wantItems)
			}
		})
	}
}
