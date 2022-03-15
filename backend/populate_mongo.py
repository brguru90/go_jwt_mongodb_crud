import multiprocessing
from unicodedata import name
import uuid,secrets,string,datetime
import pymongo
from multiprocessing import Pool
import time
 

def random_str(N):
    return ''.join(secrets.choice(string.ascii_uppercase + string.digits) for i in range(N))


count_insert=0

# just give evenly dividable numbers
def sql_populate(DATA_COUNT=1220000,BATCH_SIZE=2000):

    def insert_into_tb(i):
        try:
            myclient = pymongo.MongoClient("mongodb://root:root@localhost:27017/")
            mydb = myclient["jwt4"]
            mycol = mydb["users"]

            row={
                "name":random_str(10),
                "email":f"{random_str(5)}@{random_str(5)}.com",
                "description":random_str(400),
                "createdAt":datetime.datetime.now(),
                "updatedAt":datetime.datetime.now()
            }

            x = mycol.insert_one(row)
            # print(x.inserted_id)           
            # print(count, f"{i}th Record inserted successfully into mobile table")
            print("=",end="")
           
        except pymongo.errors.ConnectionFailure as error:
            print("Failed to connect DB", error)

        except Exception as error:
            print("Exception: ", error)

        finally:
            if myclient:
                myclient.close()

    def run_in_batch(DATA_COUNT,BATCH_SIZE):
        global count_insert
        while (DATA_COUNT:=DATA_COUNT-BATCH_SIZE)>=0:
            _process=[]
            for i in range(BATCH_SIZE):
                p=multiprocessing.Process(target=insert_into_tb,args=(i,)) 
                p.start()
                _process.append(p)
            for p in _process:
                p.join()
            count_insert+=BATCH_SIZE
            print("\n ====> Remaining",DATA_COUNT)

        print("completed",count_insert)


    # pool = Pool(processes=4)
    # pool.map(insert_into_tb, (range(DATA_COUNT)))
    run_in_batch(DATA_COUNT,BATCH_SIZE)

            

    
if __name__ == '__main__':
    start_time=time.time()
    sql_populate()
    end_time=time.time()
    print("Duration: ",end_time-start_time)